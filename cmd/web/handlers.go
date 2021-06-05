package main

import (
	"errors"
	"net/http"

	"github.com/charlesharries/pacific/pkg/data"
	"github.com/charlesharries/pacific/pkg/forms"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.tmpl", &templateData{})
}

// getNote gets the note for the given date (if any).
func (app *application) getNote(w http.ResponseWriter, r *http.Request) {
	date, err := app.readDateParam(r)
	if err != nil {
		app.notFound(w)
		return
	}

	note, err := app.models.Notes.Get(date, app.currentUser(r).ID)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.apiOK(w)
			return
		}

		app.serverError(w, err)
		return
	}

	app.apiNote(w, note)
}

func (app *application) updateNote(w http.ResponseWriter, r *http.Request) {
	date, err := app.readDateParam(r)
	if err != nil {
		app.notFound(w)
		return
	}

	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)

	form.Required("content")
	if !form.Valid() {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	note := &data.Note{
		Date:    date,
		Content: form.Get("content"),
		UserID:  app.currentUser(r).ID,
	}

	err = app.models.Notes.Upsert(note)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.apiOK(w)
}

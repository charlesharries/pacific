package main

import (
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/charlesharries/pacific/pkg/data"
	"github.com/charlesharries/pacific/pkg/forms"
	"github.com/charlesharries/pacific/pkg/models"
	"gorm.io/gorm/clause"
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
	d := r.URL.Query().Get(":date")

	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !reg.MatchString(d) {
		app.notFound(w)
		return
	}

	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		app.serverError(w, err)
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

func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []*models.Todo

	d := r.URL.Query().Get(":date")

	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !reg.MatchString(d) {
		app.notFound(w)
		return
	}

	app.gorm.Where("date = ?", d).Find(&todos)

	app.apiTodos(w, todos)
}

func (app *application) updateTodo(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get(":date")

	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !reg.MatchString(d) {
		app.notFound(w)
		return
	}

	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		app.serverError(w, err)
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

	app.gorm.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "date"}},
		DoUpdates: clause.AssignmentColumns([]string{"content"}),
	}).Create(&models.Todo{
		Date:      date,
		Completed: false,
		Content:   form.Get("content"),
		UserID:    app.currentUser(r).ID,
	})

	app.apiOK(w)
}

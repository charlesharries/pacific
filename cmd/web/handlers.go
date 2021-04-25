package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/charlesharries/pacific/pkg/forms"
	"github.com/charlesharries/pacific/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.tmpl", &templateData{})
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

	fmt.Printf("%#v", form.Get("content"))

	form.Required("content")
	if !form.Valid() {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.gorm.FirstOrCreate(&models.Note{
		Date:    date,
		Content: form.Get("content"),
		UserID:  app.currentUser(r).ID,
	})

	app.apiOK(w)
}

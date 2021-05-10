package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/charlesharries/pacific/pkg/models"
	"github.com/justinas/nosurf"
)

// serverError logs the stack trace to the errorLog and then returns a
// generic 500 error to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and description to the
// user when there's a problem with the request sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound returns a 404.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// addDefaultData adds some default data to our templates.
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.CSRFToken = nosurf.Token(r)
	// td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	td.IsAuthenticated = app.isAuthenticated(r)
	td.CacheKey = fmt.Sprint(timestamp())

	if len(os.Getenv("CACHE_KEY")) > 0 {
		td.CacheKey = os.Getenv("CACHE_KEY")
	}

	if app.session.Exists(r, "authenticatedUser") {
		td.User = app.session.Get(r, "authenticatedUser").(TemplateUser)
	}

	return td
}

// render renders a template to the response.
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

// isAuthenticated checks if the user is currently logged in.
func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}

// currentUser returns the currently logged-in user.
func (app *application) currentUser(r *http.Request) TemplateUser {
	return app.session.Get(r, "authenticatedUser").(TemplateUser)
}

func (app *application) apiOK(w http.ResponseWriter) {
	ok := map[string]interface{}{
		"error":   false,
		"message": "ok",
	}

	js, err := json.Marshal(ok)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) apiNote(w http.ResponseWriter, note *models.Note) {
	js, err := json.Marshal(note)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) apiTodos(w http.ResponseWriter, todos []*models.Todo) {
	js, err := json.Marshal(todos)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

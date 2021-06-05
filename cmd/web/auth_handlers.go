package main

import (
	"errors"
	"net/http"

	"github.com/charlesharries/pacific/pkg/data"
	"github.com/charlesharries/pacific/pkg/forms"
)

// registerForm displays the registration form.
func (app *application) registerForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.tmpl", &templateData{Form: forms.New(nil)})
}

// register creates a user in the database.
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "register.tmpl", &templateData{Form: form})
		return
	}

	// Save the user
	err = app.models.Users.Insert(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, data.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use.")
			app.render(w, r, "register.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}

		return
	}

	app.session.Put(r, "flash", "Successfully registered. Please log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// loginForm displays the login form.
func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// login handles logging the user in
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)

	id, err := app.models.Users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, data.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or password is incorrect.")
			app.render(w, r, "login.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}

		return
	}

	// Fetch the user from the DB
	user, err := app.models.Users.Get(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Save the user to the session
	app.session.Put(r, "auth.user", &TemplateUser{
		ID:    user.ID,
		Email: user.Email,
	})

	app.session.Put(r, "flash", "You've been logged in.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// logout clears the user's session and logs them out.
func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "auth.user")

	app.session.Put(r, "flash", "You've been logged out.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

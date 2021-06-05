package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// routes declares our application routes.
func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/register", dynamicMiddleware.ThenFunc(app.registerForm))
	mux.Post("/register", dynamicMiddleware.ThenFunc(app.register))
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.loginForm))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.login))
	mux.Post("/logout", dynamicMiddleware.ThenFunc(app.logout))
	mux.Get("/:date", dynamicMiddleware.ThenFunc(app.home))

	// Notes
	mux.Get("/notes/:date", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.getNote))
	mux.Post("/notes/:date", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updateNote))

	// Note routes
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}

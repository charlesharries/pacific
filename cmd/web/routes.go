package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.home)))

	// Auth routes
	mux.Get("/register", dynamicMiddleware.ThenFunc(app.registerForm))
	mux.Post("/register", dynamicMiddleware.ThenFunc(app.register))
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.loginForm))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.login))
	mux.Post("/logout", dynamicMiddleware.ThenFunc(app.logout))

	return standardMiddleware.Then(mux)
}

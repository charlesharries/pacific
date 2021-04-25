package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// middlewareGroup allows us to group a bunch of routes together under
// a chain of middleware.
type middlewareGroup struct {
	mux        *pat.PatternServeMux
	middleware alice.Chain
}

// middlewareGroupHandler is the callback passed to the middleware
// group.
type middlewareGroupHandler func(*middlewareGroup)

// routes declares our application routes.
func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()

	app.newMiddlewareGroup(mux, dynamicMiddleware, func(r *middlewareGroup) {
		r.Get("/", app.home)
		r.Get("/register", app.registerForm)
		r.Post("/register", app.register)
		r.Get("/login", app.loginForm)
		r.Post("/login", app.login)
		r.Post("/logout", app.logout)
	})

	app.newMiddlewareGroup(mux, dynamicMiddleware.Append(app.requireAuthentication), func(r *middlewareGroup) {
		r.Get("/notes/:date", app.getNote)
		r.Post("/notes/:date", app.updateNote)
	})

	// Note routes
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}

// Get passes any GET requests to this group on to the group's router.
func (c *middlewareGroup) Get(route string, h func(w http.ResponseWriter, r *http.Request)) {
	c.mux.Get(route, c.middleware.ThenFunc(h))
}

// Post passes any POST requests to this group on to the group's router.
func (c *middlewareGroup) Post(route string, h func(w http.ResponseWriter, r *http.Request)) {
	c.mux.Post(route, c.middleware.ThenFunc(h))
}

// newMiddlewareGroup creates a new middleware group.
func (app *application) newMiddlewareGroup(mux *pat.PatternServeMux, middleware alice.Chain, fn middlewareGroupHandler) {
	r := &middlewareGroup{mux, middleware}

	fn(r)
}

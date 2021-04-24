package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/charlesharries/pacific/pkg/forms"
)

type templateData struct {
	CSRFToken       string
	Form            *forms.Form
	IsAuthenticated bool
	Flash           string
	User            TemplateUser
}

// TemplateUser is the data representation of a single user rendered
// in our templates.
type TemplateUser struct {
	ID    int
	Email string
}

// human date formats time.Time objects into a human-readable string
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "pages/*.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "layouts/*.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "partials/*.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

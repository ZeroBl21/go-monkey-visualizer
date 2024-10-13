package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ZeroBl21/go-monkey-visualizer/ui"
)

// Holds the structure for any dynamic for HTML template
type templateData struct {
	CurrentYear int
}

func (app *application) newTemplateData(_ *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			// "html/partials/*.html",
			page,
		}

		ts, err := template.
			New(name).
			Funcs(functions).
			ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// "V-Table" for custom templates functions
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Returns a nicely formatted string representation of time.Time object
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 3:04pm")
}

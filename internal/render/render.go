package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/pbauer95/bookings/internal/config"
	"github.com/pbauer95/bookings/internal/models"
)

var app *config.AppConfig
var pathToTemplates = "./templates"
var functions = template.FuncMap{}

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.SessionManager.PopString(r.Context(), "flash")
	td.Warning = app.SessionManager.PopString(r.Context(), "warning")
	td.Error = app.SessionManager.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders template using html template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := templateCache[tmpl]

	if !ok {
		return errors.New("could not get template from cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buffer, td)

	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// Get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		return templateCache, err
	}

	// Range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}

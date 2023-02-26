package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/pbauer95/bookings/internal/config"
	"github.com/pbauer95/bookings/internal/models"
	"github.com/pbauer95/bookings/internal/render"
)

var app config.AppConfig
var sessionManager *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	//change this to true if in production
	app.Production = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.Production

	app.SessionManager = sessionManager

	templateCache, err := CreateTestTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplate(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)

	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.PostAvailabilityJson)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

// NoSurf adds CRSF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Production,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad saves and loads the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return app.SessionManager.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
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

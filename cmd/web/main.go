package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pbauer95/bookings/internal/config"
	"github.com/pbauer95/bookings/internal/handlers"
	"github.com/pbauer95/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

// main is the main application function
func main() {

	//change this to true if in production
	app.Production = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.Production

	app.SessionManager = sessionManager

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = serve.ListenAndServe()
}

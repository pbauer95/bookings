package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pbauer95/bookings/entities"
	"github.com/pbauer95/bookings/internal/config"
	"github.com/pbauer95/bookings/internal/handlers"
	"github.com/pbauer95/bookings/internal/helpers"
	"github.com/pbauer95/bookings/internal/render"
	"github.com/pbauer95/bookings/internal/repository"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

// main is the main application function
func main() {
	checkError(run())

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = serve.ListenAndServe()
}

func run() error {
	gob.Register(entities.Reservation{})

	//change this to true if in production
	app.Production = false

	app.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbRepo, err := Repo.InitializeDb()

	if err != nil {
		panic(err)
	}

	app.Repo = dbRepo

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.Production

	app.SessionManager = sessionManager

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	handlerRepo := handlers.NewRepo(&app)
	handlers.NewHandlers(handlerRepo)
	helpers.NewHelpers(&app)
	render.NewTemplate(&app)

	return nil
}

func checkError(err error) {
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/pbauer95/bookings/internal/repository"
)

// AppConfig holds the application config
type AppConfig struct {
	Repo           *Repo.Repo
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	Production     bool
	SessionManager *scs.SessionManager
}

package handlers

import (
	"net/http"

	"github.com/pbauer95/bookings/pkg/config"
	"github.com/pbauer95/bookings/pkg/models"
	"github.com/pbauer95/bookings/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.SessionManager.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform somr logic
	stringMap := map[string]string{
		"test": "Hello again!",
	}

	stringMap["remote_ip"] = repo.App.SessionManager.GetString(r.Context(), "remote_ip")

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

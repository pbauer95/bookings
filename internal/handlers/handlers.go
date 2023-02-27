package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pbauer95/bookings/entities"
	"github.com/pbauer95/bookings/internal/config"
	"github.com/pbauer95/bookings/internal/forms"
	"github.com/pbauer95/bookings/internal/helpers"
	"github.com/pbauer95/bookings/internal/models"
	"github.com/pbauer95/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation renders the 'make a reservation' page and displays a form
func (repo *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation entities.Reservation

	data := map[string]interface{}{
		"reservation": emptyReservation,
	}

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (repo *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := entities.Reservation{
		User: entities.User{
			FirstName: r.Form.Get("first_name"),
			LastName:  r.Form.Get("last_name"),
			Email:     r.Form.Get("email"),
		},
		Phone: sql.NullString{
			String: r.Form.Get("phone"),
			Valid:  true,
		},
	}

	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "phone")
	form.IsEmail("email")

	if !form.Valid() {
		data := map[string]interface{}{
			"reservation": reservation,
		}

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	repo.App.SessionManager.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Generals renders the room page
func (repo *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the major rooms page
func (repo *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Majors renders the search availability page
func (repo *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// Majors renders the search availability page
func (repo *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	repo.App.Repo.Connection.First(&user, 1)

	res, _ := json.Marshal(user)

	w.Write(res)
}

// Majors renders the search availability page
func (repo *Repository) PostAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func (repo *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := repo.App.SessionManager.Get(r.Context(), "reservation").(entities.Reservation)

	if !ok {
		repo.App.ErrorLog.Println("Can't get error from session")
		repo.App.SessionManager.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	repo.App.SessionManager.Remove(r.Context(), "reservation")

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{
			"reservation": reservation,
		},
	})
}

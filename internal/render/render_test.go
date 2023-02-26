package render

import (
	"net/http"
	"testing"

	"github.com/pbauer95/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()

	if err != nil {
		t.Error("lol")
		return
	}

	sessionManager.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Fatal("Session Data for Flash not found")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})

	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "non-existing.page.tmpl", &models.TemplateData{})

	if err == nil {
		t.Error("Error rendered non existing template")
	}

}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, err = sessionManager.Load(ctx, r.Header.Get("X-Session"))

	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)

	return r, nil
}

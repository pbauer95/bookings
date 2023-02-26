package render

import (
	"log"
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

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	log.Println(r)

	ctx := r.Context()
	ctx, err = sessionManager.Load(ctx, r.Header.Get("X-Session"))

	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)

	return r, nil
}

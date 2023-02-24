package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/pbauer95/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error("Type is not *chi.Mux")
	}
}

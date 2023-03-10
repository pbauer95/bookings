package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pbauer95/bookings/entities"
	"github.com/pbauer95/bookings/internal/config"
)

type myWriter struct{}

func (w *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (w *myWriter) WriteHeader(i int) {}

func (w *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

var sessionManager *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(entities.Reservation{})

	//change this to true if in production
	testApp.Production = false

	testApp.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	testApp.ErrorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false

	testApp.SessionManager = sessionManager

	app = &testApp

	os.Exit(m.Run())
}

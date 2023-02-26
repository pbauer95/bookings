package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{
		"home",
		"/",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"about",
		"/about",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"generals-quarters",
		"/generals-quarters",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"majors-suite",
		"/majors-suite",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"search-availability",
		"/search-availability",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"contact",
		"/contact",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"make-reservation",
		"/make-reservation",
		"GET",
		[]postData{},
		http.StatusOK,
	}, {
		"post-search-availability",
		"/search-availability",
		"POST",
		[]postData{
			{
				key:   "start",
				value: "2020-01-01",
			}, {
				key:   "end",
				value: "2021-01-01",
			},
		},
		http.StatusOK,
	}, {
		"post-search-availability-json",
		"/search-availability-json",
		"POST",
		[]postData{
			{
				key:   "start",
				value: "2020-01-01",
			}, {
				key:   "end",
				value: "2021-01-01",
			},
		},
		http.StatusOK,
	}, {
		"make-reservation-post",
		"/make-reservation",
		"POST",
		[]postData{
			{
				key:   "first_name",
				value: "John",
			}, {
				key:   "last_name",
				value: "Smith",
			}, {
				key:   "email",
				value: "test@test.de",
			}, {
				key:   "phone",
				value: "1234",
			},
		},
		http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range tests {

		switch test.method {
		case "GET":
			response, err := ts.Client().Get(ts.URL + test.url)

			if err != nil {
				t.Fatal(err)
				return
			}

			if response.StatusCode != test.expectedStatusCode {
				t.Fatalf("%s: Response code %d does not match expected reponse code %d", test.name, response.StatusCode, test.expectedStatusCode)
				return
			}

		case "POST":
			values := url.Values{}

			for _, x := range test.params {
				values.Add(x.key, x.value)
			}

			response, err := ts.Client().PostForm(ts.URL+test.url, values)

			if err != nil {
				t.Fatal(err)
				return
			}

			if response.StatusCode != test.expectedStatusCode {
				t.Fatalf("%s: Response code %d does not match expected reponse code %d", test.name, response.StatusCode, test.expectedStatusCode)
				return
			}

		default:
			t.Fatalf("Test %s does not use a valid HTTP method", test.name)
		}
	}
}

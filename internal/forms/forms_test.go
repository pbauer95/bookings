package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{
		"a": {"a"},
		"b": {"b"},
		"c": {"c"},
	}

	form = New(postedData)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows invalid when should be valid")
	}
}

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Has(t *testing.T) {
	form := New(url.Values{})

	if form.Has("email") {
		t.Error("form valid when should be invalid")
	}

	postedData := url.Values{
		"email": {"test@test.de"},
	}

	form = New(postedData)

	if !form.Has("email") {
		t.Error("invalid form when should be valid")
	}

}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("form is valid when it should not be")
	}

	postedData := url.Values{
		"email": {"test@"},
	}

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("form is valid when it should not be")
	}

	postedData = url.Values{
		"email": {"test@test.de"},
	}

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("showing invalid email when should be valid")
	}
}

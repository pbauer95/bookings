package forms

import "testing"

func TestErrors_Add(t *testing.T) {
	testErrors := errors{}

	testErrors.Add("test", "1")

	resp := testErrors.Get("test")

	if resp != "1" {
		t.Error("value was not correctly added to errors")
	}
}

func TestErrors_Get(t *testing.T) {
	testErrors := errors{
		"test": {"1", "2", "3"},
	}

	resp := testErrors.Get("test")

	if resp != "1" {
		t.Error("returned incorrect value")
	}

	resp = testErrors.Get("non-existing")

	if resp != "" {
		t.Error("returned value when should not")
	}
}

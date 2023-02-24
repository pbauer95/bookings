package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHandler MyHandler
	h := NoSurf(&myHandler)

	switch h.(type) {
	case http.Handler:
		return
	default:
		t.Error("Type is not http.Handler")
	}

}

func TestSessionLoad(t *testing.T) {
	var myHandler MyHandler
	h := SessionLoad(&myHandler)

	switch h.(type) {
	case http.Handler:
		return
	default:
		t.Error("Type is not http.Handler")
	}

}

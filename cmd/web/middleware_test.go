package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	// storing in v whatever h type is
	switch v := h.(type) {
	case http.Handler:
		// nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v)) // %T is type
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	// storing in v whatever h type is
	switch v := h.(type) {
	case http.Handler:
		// nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T ", v)) // %T is type
	}
}

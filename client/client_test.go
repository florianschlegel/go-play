package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallURL(t *testing.T) {
	wasCalled := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client", r.URL.Path)
		wasCalled = true
	}))
	defer ts.Close()
	CallURL(ts.URL)
	if !wasCalled {
		t.Fatal("the testserver was not called")
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestGitHubStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, nil)
	}))
	defer ts.Close()

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	c := &client{rawurl: ts.URL}
	ok, err := c.githubStatus(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("Not ok.")
	}
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitHubStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, nil)
	}))
	defer ts.Close()

	c := &client{rawurl: ts.URL}
	ok, err := c.githubStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("Not ok.")
	}
}

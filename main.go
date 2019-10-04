package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func main() {
	http.HandleFunc("/", handleGitHubStatus)
	appengine.Main()
}

func handleGitHubStatus(w http.ResponseWriter, r *http.Request) {
	c := newClient()
	ok, err := c.githubStatus(appengine.NewContext(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ok {
		fmt.Fprintln(w, "GitHub status ok!")
		return
	}

	fmt.Fprintln(w, "GitHub status ng!")
}

type client struct {
	rawurl string
}

func newClient() *client {
	return &client{rawurl: "https://github.com/"}
}

func (c *client) githubStatus(ctx context.Context) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, c.rawurl, nil)
	if err != nil {
		return false, err
	}

	resp, err := urlfetch.Client(ctx).Do(req.WithContext(ctx))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

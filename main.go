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
	ok, err := githubStatus(appengine.NewContext(r))
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

func githubStatus(ctx context.Context) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "https://github.com/", nil)
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

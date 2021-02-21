package query

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/qiyihuang/omni-cmd/internal/config"
)

// search the query using configured search engine.
func search(params []string, w http.ResponseWriter, r *http.Request) {
	baseURL := os.Getenv("SEARCH_ENGINE_URL")
	if baseURL == "" {
		baseURL = "https://www.google.com/search?q="
	}

	query := strings.Join(params, " ")
	searchURL := baseURL + url.QueryEscape(query)
	http.Redirect(w, r, searchURL, 301)
}

// handleUser handles the case when user want to go to specific user profile
// (e.g. Twitter profile)
func handleUser(arg string, qc config.Query, w http.ResponseWriter, r *http.Request) {
	username := arg[1:] // arg[0] is "@"
	rdURL := qc.URL + url.QueryEscape(username)
	http.Redirect(w, r, rdURL, 301)
}

// handleSubDir handles case when user want to go to sub url (e.g. GitHub repo)
func handleSubURL(subURL string, qc config.Query, w http.ResponseWriter, r *http.Request) {
	rdURL := qc.URL + subURL
	http.Redirect(w, r, rdURL, 301)
}

// handleSearch search in the command website.
func handleSearch(params []string, qc config.Query, w http.ResponseWriter, r *http.Request) {
	query := strings.Join(params, " ")
	rdURL := qc.URL + qc.SearchStr + url.QueryEscape(query)
	http.Redirect(w, r, rdURL, 301)
}

// Handle redirects the browser to url according to query passed.
func Handle(params []string, w http.ResponseWriter, r *http.Request) {
	if len(params) == 0 {
		return
	}

	cmd := params[0]
	queryConfig := config.QueryConfig[cmd]
	if queryConfig == (config.Query{}) {
		search(params, w, r)
		return
	}

	cmdParams := params[1:]
	if len(cmdParams) == 0 {
		http.Redirect(w, r, queryConfig.URL, 301)
		return
	}

	firstArg := cmdParams[0]
	if firstArg[0:1] == "@" {
		handleUser(firstArg, queryConfig, w, r)
		return
	}

	if subDirs := strings.Split(firstArg, "/"); len(subDirs) > 1 {
		handleSubURL(firstArg, queryConfig, w, r)
		return
	}

	handleSearch(cmdParams, queryConfig, w, r)
}

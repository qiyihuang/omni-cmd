package server

import (
	"net/http"
	"net/url"
	"strings"
)

// Mapping commands with handler functions.
var commandHandlersMap = map[string]interface{}{
	"gh": buildGitHubURL,
	"gm": buildGmailURL,
	"tw": buildTwitterURL,
	"yt": buildYouTubeURL,
}

func buildGitHubURL(params []string) string {
	baseURL := "https://www.github.com/"

	if len(params) == 0 {
		return baseURL
	}

	firstArg := params[0]
	if firstArg[0:1] == "@" {
		username := firstArg[1:]
		return baseURL + username
	}

	if repo := strings.Split(firstArg, "/"); len(repo) > 1 {
		return baseURL + firstArg
	}

	query := strings.Join(params, " ")
	return baseURL + "search?q=" + url.QueryEscape(query)

}

func buildGmailURL(params []string) string {
	baseURL := "https://www.gmail.com/"

	if len(params) == 0 {
		return baseURL
	}

	query := strings.Join(params, " ")
	return baseURL + "#search/" + url.QueryEscape(query)
}

func buildTwitterURL(params []string) string {
	baseURL := "https://www.twitter.com/"

	if len(params) == 0 {
		return baseURL
	}

	firstArg := params[0]
	if firstArg[0:1] == "@" {
		username := firstArg[1:]
		return baseURL + username
	}

	query := strings.Join(params, " ")
	return baseURL + "search?q=" + url.QueryEscape(query)
}

func buildYouTubeURL(params []string) string {
	baseURL := "https://www.youtube.com/"

	if len(params) == 0 {
		return baseURL
	}

	firstArg := params[0]
	if firstArg[0:1] == "@" {
		username := firstArg[1:]
		return baseURL + "c/" + username
	}

	query := strings.Join(params, " ")
	return baseURL + "results?search_query=" + url.QueryEscape(query)
}

func searchGoogle(query string, w http.ResponseWriter, r *http.Request) {
	searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query)
	http.Redirect(w, r, searchURL, 301)
}

// HandleGetSearch handles get request to /search route.
func handleGetSearch(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query()["cmd"][0]
	cmdSlice := strings.Split(cmd, " ")

	urlBuilder := commandHandlersMap[cmdSlice[0]]

	if urlBuilder == nil {
		searchGoogle(cmd, w, r)
		return
	}

	params := cmdSlice[1:]
	redirectURL := urlBuilder.(func([]string) string)(params)
	http.Redirect(w, r, redirectURL, 301)
}

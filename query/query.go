package query

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

// Handler returns query handler function by command string.
func Handler(cmd string) interface{} {
	return commandHandlersMap[cmd]
}

// SearchGoogle search the query using provided search engine.
func SearchGoogle(query string, w http.ResponseWriter, r *http.Request) {
	searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query)
	http.Redirect(w, r, searchURL, 301)
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

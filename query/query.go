package query

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Mapping commands with handler functions.
var commandHandlersMap = map[string]interface{}{
	"gh": handleGitHub,
	"gm": handleGmail,
	"tw": handleTwitter,
	"yt": handleYouTube,
}

// Handler returns query handler function by command string.
func Handler(cmd string) interface{} {
	return commandHandlersMap[cmd]
}

// Search search the query using provided search engine.
func Search(query string, w http.ResponseWriter, r *http.Request) {
	baseURL := os.Getenv("SEARCH_ENGINE_URL")
	if baseURL == "" {
		baseURL = "https://www.google.com/search?q="
	}

	searchURL := baseURL + url.QueryEscape(query)
	http.Redirect(w, r, searchURL, 301)
}

func handleGitHub(params []string) string {
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

func handleGmail(params []string) string {
	baseURL := "https://www.gmail.com/"

	if len(params) == 0 {
		return baseURL
	}

	query := strings.Join(params, " ")
	return baseURL + "#search/" + url.QueryEscape(query)
}

func handleTwitter(params []string) string {
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

func handleYouTube(params []string) string {
	baseURL := "https://www.youtube.com/"

	if len(params) == 0 {
		return baseURL
	}

	firstArg := params[0]
	if firstArg[0:1] == "@" {
		username := firstArg[1:]
		return baseURL + username
	}

	query := strings.Join(params, " ")
	return baseURL + "results?search_query=" + url.QueryEscape(query)
}

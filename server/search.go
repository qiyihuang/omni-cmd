package server

import (
	"net/http"
	"strings"

	"github.com/qiyihuang/omni-cmd/query"
)

// HandleGetSearch handles get request to /search route.
func handleGetSearch(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query()["cmd"][0]
	cmdSlice := strings.Split(cmd, " ")

	handler := query.Handler(cmdSlice[0])
	if handler == nil {
		query.Search(cmd, w, r)
		return
	}

	params := cmdSlice[1:]
	redirectURL := handler.(func([]string) string)(params)
	http.Redirect(w, r, redirectURL, 301)
}

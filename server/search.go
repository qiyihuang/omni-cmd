package server

import (
	"net/http"
	"strings"

	"github.com/qiyihuang/omni-cmd/query"
)

// HandleGetSearch handles get request to /search route.
func handleGetSearch(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()["cmd"][0]
	params := strings.Split(qs, " ")

	query.Handle(params, w, r)
}

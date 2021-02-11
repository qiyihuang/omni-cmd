package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/search", handleGetSearch).Methods("GET")

	return muxRouter
}

// Run starts http server.
func Run() error {
	mux := makeMuxRouter()
	srv := &http.Server{
		Addr:           ":" + os.Getenv("Port"),
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

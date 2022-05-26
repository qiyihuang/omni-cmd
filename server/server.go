package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/qiyihuang/omni-cmd/metrics"
)

// Start starts http server.
func Start() error {
	muxRouter := mux.NewRouter()
	muxRouter.Use(logging)
	muxRouter.HandleFunc("/search", handleGetSearch).Methods("GET")

	srv := &http.Server{
		Addr:           "0.0.0.0:" + os.Getenv("PORT"),
		Handler:        muxRouter,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := metrics.StartTimer("request")
		next.ServeHTTP(w, r)
		t.End()
	})
}
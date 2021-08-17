package main

import (
	"net/http"

	"github.com/NamanBalaji/mise/internal/handlers"
	"github.com/go-chi/chi"
)

// routes returns a http Handler that routes requests based on URL
func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/ping", handlers.Repo.Ping)

	mux.Post("/set", handlers.Repo.Set)

	mux.Post("/get", handlers.Repo.Get)

	mux.Post("/get-range", handlers.Repo.GetRange)
	return mux
}

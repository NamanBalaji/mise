package main

import (
	"net/http"

	"github.com/NamanBalaji/mise/internal/handlers"
	"github.com/go-chi/chi"
)

// routes returns a http Handler that routes requests based on URL
func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)

	mux.Get("/ping", handlers.Repo.Ping)

	return mux
}

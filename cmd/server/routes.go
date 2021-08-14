package main

import (
	"net/http"

	"github.com/NamanBalaji/mise/internal/handlers"
	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)

	mux.Get("/ping", handlers.Repo.Ping)

	return mux
}

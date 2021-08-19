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

	mux.Post("/set-list", handlers.Repo.SetList)

	mux.Post("/get", handlers.Repo.Get)

	mux.Post("/get-list", handlers.Repo.GetListNodeValue)

	mux.Delete("/delete", handlers.Repo.Delete)

	mux.Post("/get-range", handlers.Repo.GetRange)

	mux.Post("/add", handlers.Repo.Add)

	mux.Post("/add-list", handlers.Repo.AddToLinkedList)

	mux.Delete("/delete-element", handlers.Repo.DeleteIndex)

	return mux
}

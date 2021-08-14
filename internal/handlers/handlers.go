package handlers

import (
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/database"
)

type Repository struct {
	App *config.AppConfig
	DB  *database.DB
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

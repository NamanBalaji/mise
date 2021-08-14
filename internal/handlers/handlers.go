package handlers

import (
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/database"
)

// Repository is a struct that contains  pointers to DB and App objects
type Repository struct {
	App *config.AppConfig
	DB  *database.DB
}

var Repo *Repository

// NewRepo returns a pointer to the Repository
func NewRepo(a *config.AppConfig, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

//Ping handles requests directed to /ping and responds with pong if server is up
func (m *Repository) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

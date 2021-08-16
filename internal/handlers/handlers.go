package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/database"
	"github.com/NamanBalaji/mise/internal/resp"
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

// Set is used to set a key with a value
func (m *Repository) Set(w http.ResponseWriter, r *http.Request) {
	var body resp.SetRequest
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error occurred while reading the request body"))
	}
	err = json.Unmarshal(req, &body)
	if err != nil {
		w.Write([]byte("please check if the request body hs the correct structure"))
	}
	response, err := m.DB.Set(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}

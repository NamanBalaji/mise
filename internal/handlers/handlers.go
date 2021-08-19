package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/database"
	"github.com/NamanBalaji/mise/internal/helpers"
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

	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
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

// Get is used to get a key's value
func (m *Repository) Get(w http.ResponseWriter, r *http.Request) {
	var body resp.GetRequest

	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}

	response, err := m.DB.Get(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}

func (m *Repository) Delete(w http.ResponseWriter, r *http.Request) {
	var body resp.DeleteRequest
	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.Delete(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// GetRange is the handler for getting a portion of an array
func (m *Repository) GetRange(w http.ResponseWriter, r *http.Request) {
	var body resp.GetRangeRequest
	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.GetRange(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// Add is the handler for appending an element into an array
func (m *Repository) Add(w http.ResponseWriter, r *http.Request) {
	var body resp.AddToArrayRequest
	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.AddToArray(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// DeleteIndex is the handler for deleting an element from an array
func (m *Repository) DeleteIndex(w http.ResponseWriter, r *http.Request) {
	var body resp.DeleteFromArrayRequest
	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.DeleteFromArray(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// SetList is the handler for setting a key with a linked list
func (m *Repository) SetList(w http.ResponseWriter, r *http.Request) {
	var body resp.SetRequest

	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}

	response, err := m.DB.SetList(&body)

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

func (m *Repository) GetListNodeValue(w http.ResponseWriter, r *http.Request) {
	var body resp.GetListNodeRequest

	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}

	response, err := m.DB.GetListNodeValue(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}

// AddToLinkedList is the handler for appending a value to linked list
func (m *Repository) AddToLinkedList(w http.ResponseWriter, r *http.Request) {
	var body resp.AddToListRequest
	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.AddToLinkedList(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}

func (m *Repository) DeleteFromLinkedList(w http.ResponseWriter, r *http.Request) {
	var body resp.DeleteListNodeRequest

	err := helpers.RequestToStruct(&body, r)
	if err != nil {
		w.Write([]byte(`{
			"Error": "error occurred while reading the request body, please check if the request body hs the correct structure"
		}`))
		return
	}
	response, err := m.DB.DeleteFromLinkedList(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(resp)
}

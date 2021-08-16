package database

import (
	"strings"
	"sync"

	"github.com/NamanBalaji/mise/internal/resp"
)

// DB is a struct that holds are database and it's characteristics
type DB struct {
	database map[string]interface{}
	memory   bool
	mu       sync.Mutex
}

// NewDB returns a pointer to the DB
func NewDB(memory bool) *DB {
	return &DB{
		database: make(map[string]interface{}),
		memory:   memory,
	}
}

// Set is used to set a key in database
func (db *DB) Set(r *resp.SetRequest) (resp.SetResponse, error) {

	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse
	if r.Key == "" || r.Value == nil {
		response.Status = -1
		response.Message = "please provide valid key and value"
		return response, nil
	}

	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		response.Message = "key already present please use UPDATE if you are trying to update"
		response.Status = 1
		return response, nil
	}

	db.database[strings.ToLower(r.Key)] = r.Value

	response.Message = "OK"
	response.Status = 0
	return response, nil
}

// Get is used to get a value associated with a key
func (db *DB) Get(r *resp.GetRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	response := resp.GetResponse{
		Message: "No such key present",
		Value:   nil,
		Status:  1,
	}

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		response.Value = val
		response.Message = "OK"
		response.Status = 0
	}

	return response, nil
}

package database

import (
	"sync"
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

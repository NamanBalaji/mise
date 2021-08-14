package database

import (
	"sync"
)

type DB struct {
	database map[string]interface{}
	memory   bool
	mu       sync.Mutex
}

func NewDB(memory bool) *DB {
	return &DB{
		database: make(map[string]interface{}),
		memory:   memory,
	}
}

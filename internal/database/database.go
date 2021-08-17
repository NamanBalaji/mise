package database

import (
	"fmt"
	"reflect"
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

// GetRange returns a given portion of a slice
func (db *DB) GetRange(r *resp.GetRangeRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	response := resp.GetResponse{
		Message: "No such key present",
		Value:   nil,
		Status:  1,
	}

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(val)) != "[]interface {}" {
			response.Message = "the associated value is not an array"
			return response, nil
		}
		value := val.([]interface{})
		if r.Start < 0 || r.Stop > len(value) {
			response.Message = "range does not exist"
			return response, nil
		}
		if r.Stop == -1 {
			r.Stop = len(value)
		}
		response.Value = value[r.Start:r.Stop]
		response.Message = "OK"
		response.Status = 0
	}
	return response, nil
}

// AddToArray appends at a particular index or at the end
func (db *DB) AddToArray(r *resp.AddToArrayRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	response := resp.SetResponse{
		Message: "No such key present",
		Status:  1,
	}

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(val)) != "[]interface {}" {
			response.Message = "the associated value is not an array"
			return response, nil
		}
		value := val.([]interface{})
		if r.Index < -1 || r.Index >= len(value) {
			response.Message = "given index does not exist"
			response.Status = 1
			return response, nil
		}
		if r.Index == -1 {
			value = append(value, r.Value)
			db.database[strings.ToLower(r.Key)] = value
			response.Message = "OK"
			response.Status = 0
		} else {

			var modifiedArray []interface{}
			modifiedArray = append(modifiedArray, value[:r.Index]...)
			modifiedArray = append(modifiedArray, r.Value)
			modifiedArray = append(modifiedArray, value[r.Index:]...)
			db.database[strings.ToLower(r.Key)] = modifiedArray
			response.Message = "OK"
			response.Status = 0
		}
	}
	return response, nil
}

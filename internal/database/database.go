package database

import (
	"errors"
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
		return response, errors.New("invalid key or value")
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

	var response resp.GetResponse

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		response.Value = val
		response.Message = "OK"
		response.Status = 0

		return response, nil
	}
	return response, errors.New("no such key present")
}

// GetRange returns a given portion of a slice
func (db *DB) GetRange(r *resp.GetRangeRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.GetResponse

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(val)) != "[]interface {}" {
			return response, errors.New("associated value is not an array")
		}
		value := val.([]interface{})
		if r.Start < 0 || r.Stop > len(value) {
			return response, errors.New("range does not exist")
		}
		if r.Stop == -1 {
			r.Stop = len(value)
		}
		response.Value = value[r.Start:r.Stop]
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}
	return response, errors.New("no such key present")
}

// AddToArray appends at a particular index or at the end
func (db *DB) AddToArray(r *resp.AddToArrayRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(val)) != "[]interface {}" {
			return response, errors.New("associated value is not an array")
		}
		value := val.([]interface{})
		if r.Index < -1 || r.Index >= len(value) {
			return response, errors.New("given index does not exist")
		}
		if r.Index == -1 {
			value = append(value, r.Value)
			db.database[strings.ToLower(r.Key)] = value
			response.Message = "OK"
			response.Status = 0
			return response, nil
		} else {
			var modifiedArray []interface{}
			modifiedArray = append(modifiedArray, value[:r.Index]...)
			modifiedArray = append(modifiedArray, r.Value)
			modifiedArray = append(modifiedArray, value[r.Index:]...)
			db.database[strings.ToLower(r.Key)] = modifiedArray
			response.Message = "OK"
			response.Status = 0
			return response, nil
		}
	}
	return response, errors.New("no such key present")
}

// DeleteFromArray deletes the element at the given index
func (db *DB) DeleteFromArray(r *resp.DeleteFromArrayRequest) (resp.DeleteResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.DeleteResponse

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(val)) != "[]interface {}" {
			return response, errors.New("associated value is not an array")
		}
		value := val.([]interface{})
		if r.Index < -1 || r.Index >= len(value) {
			return response, errors.New("given index does not exist")
		}

		if r.Index == 0 {
			response.Value = value[0]
			value = value[1:]
			db.database[strings.ToLower(r.Key)] = value
			response.Message = "OK"
			response.Status = 0
			return response, nil
		} else if r.Index == len(value)-1 {
			response.Value = value[r.Index]
			value = value[:len(value)-1]
			db.database[strings.ToLower(r.Key)] = value
			response.Message = "OK"
			response.Status = 0
			return response, nil
		} else {
			response.Value = value[r.Index]
			var modifiedArray []interface{}
			modifiedArray = append(modifiedArray, value[:r.Index]...)
			modifiedArray = append(modifiedArray, value[r.Index+1:]...)
			db.database[strings.ToLower(r.Key)] = modifiedArray
			response.Message = "OK"
			response.Status = 0
			return response, nil
		}
	}
	return response, errors.New("no such key present")
}

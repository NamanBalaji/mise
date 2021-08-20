package database

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/NamanBalaji/mise/internal/dataStructures"
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

// Delete deletes a key from the database
func (db *DB) Delete(r *resp.DeleteRequest) (resp.DeleteResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.DeleteResponse

	if val, ok := db.database[strings.ToLower(r.Key)]; ok {
		delete(db.database, r.Key)
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
		} else if r.Index == -1 {
			response.Value = value[len(value)-1]
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

// SetList is used to set a linked list
func (db *DB) SetList(r *resp.SetRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse
	if r.Key == "" || r.Value == nil {
		return response, errors.New("invalid key or value")
	}

	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		response.Message = "key already present please use ADD if you are trying to add to the list"
		response.Status = 1
		return response, nil
	}

	list := dataStructures.NewLinkedList()

	// if the value is an array create a new list with i nodes
	if fmt.Sprint(reflect.TypeOf(r.Value)) == "[]interface {}" {
		array := r.Value.([]interface{})
		for _, v := range array {
			list.AddTail(v)
		}
		db.database[r.Key] = list
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}

	// if only a single value passed
	list.AddHead(r.Value)
	db.database[r.Key] = list
	response.Message = "OK"
	response.Status = 0
	return response, nil
}

// GetListNodeValue returns the value of the first or last node of the linked list
func (db *DB) GetListNodeValue(r *resp.GetListNodeRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.GetResponse

	if list, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(list)) != "*dataStructures.LinkedList" {
			return response, errors.New("associated value is not of type list")
		}
		linkedList := list.(*dataStructures.LinkedList)
		if r.GetFirst {
			response.Value = linkedList.GetFirst().Value()
		} else {
			response.Value = linkedList.GetLast().Value()
		}
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}
	return response, errors.New("no such key present")
}

// AddToLinkedList adds a node to the start or end of the linked list
func (db *DB) AddToLinkedList(r *resp.AddToListRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse

	if list, ok := db.database[strings.ToLower(r.Key)]; ok {

		if fmt.Sprint(reflect.TypeOf(list)) != "*dataStructures.LinkedList" {
			return response, errors.New("associated value is not of type list")
		}

		linkedList := db.database[r.Key].(*dataStructures.LinkedList)
		if r.AddFirst {
			linkedList.AddHead(r.Value)
		} else {
			linkedList.AddTail(r.Value)
		}
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}
	return response, errors.New("no such key present")
}

//DeleteFromLinkedList deletes the first or last node from the linked list
func (db *DB) DeleteFromLinkedList(r *resp.DeleteListNodeRequest) (resp.DeleteResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.DeleteResponse
	if list, ok := db.database[strings.ToLower(r.Key)]; ok {
		if fmt.Sprint(reflect.TypeOf(list)) != "*dataStructures.LinkedList" {
			return response, errors.New("associated value is not of type list")
		}

		linkedList := db.database[r.Key].(*dataStructures.LinkedList)
		if r.DelFirst {
			response.Value = linkedList.DelHead().Value()
		} else {
			response.Value = linkedList.DelTail().Value()
		}
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}
	return response, errors.New("no such key present")
}

//SetSortedSet initializes a sorted set and stores it into the DB
func (db *DB) SetSortedSet(r *resp.SetRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse
	if r.Key == "" || r.Value == nil {
		return response, errors.New("invalid key or value")
	}

	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		response.Message = "key already present please use ADD if you are trying to add to the list"
		response.Status = 1
		return response, nil
	}

	var sSet *dataStructures.SortedSet

	fmt.Println(reflect.TypeOf(r.Value))
	if fmt.Sprint(reflect.TypeOf(r.Value)) == "[]interface {}" || fmt.Sprint(reflect.TypeOf(r.Value)) == "float64" {
		if fmt.Sprint(reflect.TypeOf(r.Value)) == "[]interface {}" {
			for i, v := range r.Value.([]interface{}) {
				if fmt.Sprint(reflect.TypeOf(v)) != "float64" {
					return response, errors.New("sorted set only takes in integer values")
				}
				if i == 0 {
					sSet = dataStructures.NewSortedSet([]float64{v.(float64)})
				} else {
					sSet.Add(v.(float64))
				}

			}
		} else {
			sSet = dataStructures.NewSortedSet([]float64{r.Value.(float64)})
		}
	} else {
		return response, errors.New("sorted set only takes in integer values")
	}

	db.database[r.Key] = sSet
	response.Message = "OK"
	response.Status = 0
	return response, nil
}

//AddSortedSet adds a value to sorted set
func (db *DB) AddSortedSet(r *resp.SetRequest) (resp.SetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.SetResponse
	if r.Key == "" || r.Value == nil {
		return response, errors.New("invalid key or value")
	}
	if fmt.Sprint(reflect.TypeOf(r.Value)) != "float64" {
		return response, errors.New("sorted set oly takesin integer values")
	}
	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		sSet := db.database[r.Key].(*dataStructures.SortedSet)
		sSet.Add(r.Value.(float64))
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}

	return response, errors.New("no such key present")
}

// GetFromSortedSet retrieves the min or max value from the sorted set
func (db *DB) GetFromSortedSet(r *resp.SSetGDRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.GetResponse
	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		sSet := db.database[r.Key].(*dataStructures.SortedSet)

		if sSet.Size() == 0 {
			return response, errors.New("set is empty")
		}
		if r.Max {
			response.Value = sSet.GetMax()
		} else {
			response.Value = sSet.GetMin()
		}
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}

	return response, errors.New("no such key present")
}

// DeleteFromSortedSet deletes the min or max value from the sorted set
func (db *DB) DeleteFromSortedSet(r *resp.SSetGDRequest) (resp.GetResponse, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var response resp.GetResponse
	if _, ok := db.database[strings.ToLower(r.Key)]; ok {
		sSet := db.database[r.Key].(*dataStructures.SortedSet)
		if sSet.Size() == 0 {
			return response, errors.New("set is empty")
		}
		if r.Max {
			response.Value = sSet.DeleteMax()
		} else {
			response.Value = sSet.DeleteMin()
		}
		response.Message = "OK"
		response.Status = 0
		return response, nil
	}

	return response, errors.New("no such key present")
}

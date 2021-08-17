package resp

// request structure for setting a key
type SetRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// request structure for retrieving a key
type GetRequest struct {
	Key string `json:"key"`
}

// request structure for retrieving a portion of an array
type GetRangeRequest struct {
	Key   string `json:"key"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
}

// request structure for adding an element to an array
type AddToArrayRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Index int         `json:"index"`
}

// request structure for deleteing a key-value pair
type DeleteRequest struct {
	Key string `json:"key"`
}

// request structure for deleting an element from an array
type DeleteFromArrayRequest struct {
	Key   string `json:"key"`
	Index int    `json:"index"`
}

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

// request structure for adding to a list
type AddToListRequest struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	AddFirst bool        `json:"add_first"`
}

// request structure for getting a node value
type GetListNodeRequest struct {
	Key      string `json:"key"`
	GetFirst bool   `json:"get_first"`
}

// request structure for deleteing a linked list node
type DeleteListNodeRequest struct {
	Key      string `json:"key"`
	DelFirst bool   `json:"delete_first"`
}

type SSetGDRequest struct {
	Key string `json:"key"`
	Max bool   `json:"max"`
}

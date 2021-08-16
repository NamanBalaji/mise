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

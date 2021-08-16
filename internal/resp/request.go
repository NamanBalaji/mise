package resp

// request structure for setting a key
type SetRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// requesting structure for retrieving a key
type GetRequest struct {
	Key string `json:"key"`
}

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

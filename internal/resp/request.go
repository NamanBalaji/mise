package resp

// request structure for setting a key
type SetRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

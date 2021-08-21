package resp

// response returned when a key is set
type SetResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// response structure when a key's value if fetched
type GetResponse struct {
	Value   interface{} `json:"value"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

// response structure when a delete action completes
type DeleteResponse struct {
	Value   interface{} `json:"value"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

// response structure when size of a key is requested
type SizeResponse struct {
	Value int `json:"value"`
}

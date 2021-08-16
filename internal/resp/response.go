package resp

type SetResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type GetResponse struct {
	Value   interface{} `json:"value"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

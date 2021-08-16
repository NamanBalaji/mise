package resp

type SetResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

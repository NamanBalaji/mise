package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestToStruct marshals request body to desired struct
func RequestToStruct(body interface{}, req *http.Request) error {
	r, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(r, body)
	if err != nil {
		return err
	}

	return nil
}

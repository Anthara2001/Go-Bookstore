package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(req *http.Request, x interface{}) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}
	return nil

}

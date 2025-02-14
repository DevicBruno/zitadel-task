package server

import (
	"encoding/json"
	"net/http"
)

// Response represents the API response structure.
type Response struct {
	Message string `json:"message"`
}

// jsonResponse writes a JSON response to the response writer.
func jsonResponse(w http.ResponseWriter, resp Response, status int) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

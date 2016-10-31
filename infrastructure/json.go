package infrastructure

import (
	"encoding/json"
	"net/http"
)

func WriteJson(status int, response interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

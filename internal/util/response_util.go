package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
    Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(v)
}

func JSONError(w http.ResponseWriter, status int, err string) {
    JSON(w, status, ErrorResponse{Error: err})
}

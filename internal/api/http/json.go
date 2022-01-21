package http

import (
	"encoding/json"
	"net/http"
)

func toJsonResponse(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

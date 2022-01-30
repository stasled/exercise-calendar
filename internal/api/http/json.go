package http

import (
	"encoding/json"
	"net/http"
	"time"
)

func toJsonResponse(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func jsonToDatetime(dateTime string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	res, err := time.Parse(layout, dateTime)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}

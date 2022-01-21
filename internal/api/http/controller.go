package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (e Events) GetAll(rw http.ResponseWriter, r *http.Request) {
	events, err := e.service.GetEvents(e.ctx)
	if err != nil {
		http.Error(rw, "Unable to get events", http.StatusInternalServerError)
		return
	}

	toJsonResponse(rw, events)
}

func (e Events) UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	event, err := e.service.GetEventByID(e.ctx, newEvent.Id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%v event with ID: %s: ", err, newEvent.Id), http.StatusInternalServerError)
		return
	}

	err = e.service.UpdateEvent(e.ctx, event, newEvent.Title, newEvent.StartAt, newEvent.EndAt)
	if err != nil {
		http.Error(rw, fmt.Sprintf("unable to update event. %v", err), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Event successfully updated"

	toJsonResponse(rw, resp)
}

func (e Events) AddEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	err = e.service.CreateEvent(e.ctx, newEvent.Title, newEvent.StartAt, newEvent.EndAt)
	if err != nil {
		http.Error(rw, fmt.Sprintf("unable to create event. %v", err), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Event successfully created"

	toJsonResponse(rw, resp)
}

func (e Events) DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	_, err = e.service.GetEventByID(e.ctx, newEvent.Id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%v event with ID: %s: ", err, newEvent.Id), http.StatusInternalServerError)
		return
	}

	err = e.service.DeleteEvent(e.ctx, newEvent.Id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("unable to delete event. %v", err), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Event successfully deleted"

	toJsonResponse(rw, resp)
}

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"mycalendar/internal/api"
	"net/http"
	"time"
)

type events struct {
	ctx     context.Context
	service api.Service
	logger  *zap.Logger
}

func NewController(ctx context.Context, s api.Service, l *zap.Logger) *events {
	return &events{
		ctx:     ctx,
		service: s,
		logger:  l,
	}
}

type Event struct {
	Id      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	StartAt string `json:"start_at,omitempty"`
	EndAt   string `json:"end_at,omitempty"`
}

// ShowAccount GetAll
// @Summary    Gett all events
// @Tags 	   rest
// @Accept     json
// @Produce    json
// @Success    200 {array} Event
// @Failure    500 {string} string "Unable to get events"
// @Router     / [get]
func (e events) GetAll(rw http.ResponseWriter, _ *http.Request) {
	events, err := e.service.GetEvents(e.ctx)
	if err != nil {
		http.Error(rw, "Unable to get events", http.StatusInternalServerError)
		return
	}
	toJsonResponse(rw, events)
}

// ShowAccount  UpdateEvent
// @Summary 	Update event
// @Tags 		rest
// @Accept      json
// @Produce     json
// @Param 		id 	     query string false "Id"
// @Param 		title    query string false	"Title"
// @Param 		start_at query string false	"StartAt" Format(dateTime)
// @Param 		end_at   query string false	"EndAt"   Format(dateTime)
// @Success     200 {}
// @Failure     404 {string} string "not found"
// @Failure     500 {string} string
// @Router      / [put]
func (e events) UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json from request", http.StatusInternalServerError)
		return
	}

	event, err := e.service.GetEventByID(e.ctx, newEvent.Id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%v event with ID: %s: ", err, newEvent.Id), http.StatusNotFound)
		return
	}

	var startAt, endAt time.Time
	if newEvent.StartAt != "" {
		startAt, err = jsonToDatetime(newEvent.StartAt)
		if err != nil {
			http.Error(rw, "Unable to unmarshal start at from request", http.StatusInternalServerError)
			return
		}
	}

	if newEvent.StartAt != "" {
		endAt, err = jsonToDatetime(newEvent.EndAt)
		if err != nil {
			http.Error(rw, "Unable to unmarshal end at from request", http.StatusInternalServerError)
			return
		}
	}

	err = e.service.UpdateEvent(e.ctx, event, newEvent.Title, startAt, endAt)
	if err != nil {
		http.Error(rw, fmt.Sprintf("unable to update event. %v", err), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Event successfully updated"

	toJsonResponse(rw, resp)
}

// ShowAccount  AddEvent
// @Summary 	Create event
// @Tags 		rest
// @Accept      json
// @Produce     json
// @Param 		title    query string true	"Title"
// @Param 		start_at query string true	"StartAt" Format(dateTime)
// @Param 		end_at   query string true	"EndAt"   Format(dateTime)
// @Success     200 {}
// @Failure     500 {string} string
// @Router      / [post]
func (e events) AddEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json from request", http.StatusInternalServerError)
		return
	}

	var startAt, endAt time.Time
	if newEvent.StartAt != "" {
		startAt, err = jsonToDatetime(newEvent.StartAt)
		if err != nil {
			http.Error(rw, "Unable to unmarshal start at from request", http.StatusInternalServerError)
			return
		}
	}

	if newEvent.StartAt != "" {
		endAt, err = jsonToDatetime(newEvent.EndAt)
		if err != nil {
			http.Error(rw, "Unable to unmarshal end at from request", http.StatusInternalServerError)
			return
		}
	}

	err = e.service.CreateEvent(e.ctx, newEvent.Title, startAt, endAt)
	if err != nil {
		http.Error(rw, fmt.Sprintf("unable to create event. %v", err), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Event successfully created"

	toJsonResponse(rw, resp)
}

// ShowAccount  DeleteEvent
// @Summary 	Delete event
// @Tags 		rest
// @Accept      json
// @Produce     json
// @Param 		id  query string true "Id"
// @Success     200 {string}
// @Failure     404 {string} string "not found"
// @Failure     500 {string} string
// @Router      / [put]
func (e events) DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json from request", http.StatusInternalServerError)
		return
	}

	_, err = e.service.GetEventByID(e.ctx, newEvent.Id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%v event with ID: %s: ", err, newEvent.Id), http.StatusNotFound)
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

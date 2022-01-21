package proto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"mycalendar/internal/api"
	"net"
)

func RunServer(ctx context.Context, service api.Service, host, port string) {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen %s port: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	RegisterEventerServer(grpcServer, &server{
		ctx:     ctx,
		service: service,
	})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("unable to run grpc server: %v", err)
	}
}

type server struct {
	ctx     context.Context
	service api.Service
}

func (s server) CreateEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	err := s.service.CreateEvent(ctx, event.Title, event.StartAt, event.EndAt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to create event: %v", err))
	}
	return nil, nil
}

func (s server) UpdateEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	e, err := s.service.GetEventByID(ctx, event.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to find event: %v", err))
	}

	err = s.service.UpdateEvent(ctx, e, event.Title, event.StartAt, event.EndAt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to update event: %v", err))
	}
	return nil, nil
}

func (s server) DeleteEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	err := s.service.DeleteEvent(ctx, event.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to delete event: %v", err))
	}
	return nil, nil
}

func (s server) GetEventByID(ctx context.Context, event *Event) (*Event, error) {
	e, err := s.service.GetEventByID(ctx, event.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to find event: %v", err))
	}
	return &Event{Id: e.Id, Title: e.Title, StartAt: e.StartAt, EndAt: e.EndAt}, nil
}

func (s server) GetEvents(ctx context.Context, _ *EventsRequest) (*EventsResponse, error) {
	events, err := s.service.GetEvents(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to get events: %v", err))
	}

	er := make([]*Event, len(events))
	for _, e := range events {
		er = append(er, &Event{
			Id:      e.Id,
			Title:   e.Title,
			StartAt: e.StartAt,
			EndAt:   e.EndAt,
		})
	}

	return &EventsResponse{Events: er}, nil
}

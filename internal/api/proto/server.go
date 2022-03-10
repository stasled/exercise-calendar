package proto

import (
	"context"
	"fmt"
	"mycalendar/internal/api"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func RunServer(ctx context.Context, service api.Service, logger *zap.Logger) {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d",
			viper.GetString("GRPC_SERVER_HOST"),
			viper.GetInt("GRPC_SERVER_PORT")))
	if err != nil {
		logger.Fatal("failed to listen port: ",
			zap.Int("port", viper.GetInt("GRPC_SERVER_PORT")),
			zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	RegisterEventerServer(grpcServer, &server{
		ctx:     ctx,
		service: service,
	})
	logger.Info("run grpc server")
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("unable to run grpc server.", zap.Error(err))
	}
}

type server struct {
	ctx     context.Context
	service api.Service
}

func (s server) CreateEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	err := s.service.CreateEvent(ctx, event.Title, event.StartAt.AsTime(), event.EndAt.AsTime())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to create event: %v", err))
	}
	return nil, nil
}

func (s server) UpdateEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	e, err := s.service.GetEventByID(ctx, int(event.Id))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to find event: %v", err))
	}

	err = s.service.UpdateEvent(ctx, e, event.Title, event.StartAt.AsTime(), event.EndAt.AsTime())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to update event: %v", err))
	}
	return nil, nil
}

func (s server) DeleteEvent(ctx context.Context, event *Event) (*EventResponse, error) {
	err := s.service.DeleteEvent(ctx, int(event.Id))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to delete event: %v", err))
	}
	return nil, nil
}

func (s server) GetEventByID(ctx context.Context, event *Event) (*Event, error) {
	e, err := s.service.GetEventByID(ctx, int(event.Id))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to find event: %v", err))
	}
	return &Event{
		Id:      int64(e.Id),
		Title:   e.Title,
		StartAt: timestamppb.New(e.StartAt),
		EndAt:   timestamppb.New(e.EndAt),
	}, nil
}

func (s server) GetEvents(ctx context.Context, _ *EventsRequest) (*EventsResponse, error) {
	events, err := s.service.GetEvents(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Unable to get events: %v", err))
	}

	er := make([]*Event, len(events))
	for _, e := range events {
		er = append(er, &Event{
			Id:      int64(e.Id),
			Title:   e.Title,
			StartAt: timestamppb.New(e.StartAt),
			EndAt:   timestamppb.New(e.EndAt),
		})
	}

	return &EventsResponse{Events: er}, nil
}

package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"mycalendar/internal/entity"
)

var ErrSQLQuery = errors.New("sql query failed")

type repository struct {
	client *pgxpool.Pool
	logger *zap.Logger
}

func NewStorage(client *pgxpool.Pool, logger *zap.Logger) *repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Add(ctx context.Context, e entity.Event) error {
	query := "INSERT INTO events(title, start_at, end_at) VALUES($1, $2, $3)"

	r.logger.Debug(
		"SQL Query: ",
		zap.String("query", query),
		zap.String("title", e.Title),
		zap.Time("start_at", e.StartAt),
		zap.Time("end_at", e.EndAt),
	)
	_, err := r.client.Exec(ctx, query, e.Title, e.StartAt, e.EndAt)
	if err != nil {
		r.logger.Error(fmt.Sprintf("QueryRow to create event failed: %v\n", err))
		return ErrSQLQuery
	}
	return nil
}

func (r *repository) Update(ctx context.Context, e entity.Event) error {
	query := "UPDATE events SET title=$1, start_at=$2, end_at=$3 WHERE id=$4"

	r.logger.Debug(
		"SQL Query: ",
		zap.String("query", query),
		zap.String("title", e.Title),
		zap.Time("start_at", e.StartAt),
		zap.Time("end_at", e.EndAt),
		zap.Int("id", e.Id),
	)
	res, err := r.client.Exec(ctx, query, e.Title, e.StartAt, e.EndAt, e.Id)
	if err != nil || res.RowsAffected() != 1 {
		r.logger.Error(fmt.Sprintf("Query to update event failed: %v\n", err))
		return ErrSQLQuery
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM events WHERE id=$1"

	r.logger.Debug(
		"SQL Query: ",
		zap.String("query", query),
		zap.Int("id", id),
	)
	_, err := r.client.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Query to delete event failed: %v\n", err))
		return ErrSQLQuery
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id int) (entity.Event, error) {
	query := "SELECT id, title, start_at, end_at FROM events WHERE id=$1"

	r.logger.Debug(
		"SQL Query: ",
		zap.String("query", query),
		zap.Int("id", id),
	)
	var e entity.Event
	err := r.client.QueryRow(ctx, query, id).Scan(&e.Id, &e.Title, &e.StartAt, &e.EndAt)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Query to get event failed: %v\n", err))
		return entity.Event{}, ErrSQLQuery
	}
	return e, nil
}

func (r *repository) GetAll(ctx context.Context) (map[int]entity.Event, error) {
	query := "SELECT id, title, start_at, end_at FROM events"

	r.logger.Debug(
		"SQL Query: ",
		zap.String("query", query),
	)
	events := make(map[int]entity.Event, 0)
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Query to get event failed: %v\n", err))
		return nil, ErrSQLQuery
	}

	for rows.Next() {
		var e entity.Event
		err = rows.Scan(&e.Id, &e.Title, &e.StartAt, &e.EndAt)
		if err != nil {
			return nil, err
		}
		events[e.Id] = e
	}
	return events, nil
}

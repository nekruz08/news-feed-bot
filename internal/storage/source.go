package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/nekruz08/news-feed-bot/internal/model"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

func NewSourceStorage(db *sqlx.DB) *SourcePostgresStorage {
	return &SourcePostgresStorage{db: db}
}

func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var sources []model.Source
	if err := conn.SelectContext(ctx, &sources, `SELECT * FROM sources`); err != nil {
		return nil, err
	}

	return sources, nil
}

var ErrSourceNotFound = errors.New("source not found")

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var source model.Source
	if err := conn.GetContext(ctx, &source, `SELECT * FROM sources WHERE id = $1`, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrSourceNotFound // Возвращаем специальную ошибку
		}
		return nil, err
	}

	return &source, nil
}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int64

	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO sources (name, feed_url, priority)
					VALUES ($1, $2, $3) RETURNING id;`,
		source.Name, source.FeedURL, source.Priority,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SourcePostgresStorage) SetPriority(ctx context.Context, id int64, priority int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE sources SET priority = $1 WHERE id = $2`, priority, id)

	return err
}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, `DELETE FROM sources WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}

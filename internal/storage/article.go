package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/nekruz08/news-feed-bot/internal/model"
)

type ArticlePostgresStorage struct {
	db *sqlx.DB
}

func NewArticleStorage(db *sqlx.DB) *ArticlePostgresStorage {
	return &ArticlePostgresStorage{db: db}
}

func (s *ArticlePostgresStorage) Store(ctx context.Context, article model.Article) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`INSERT INTO articles (source_id, title, link, published_at)
	    				VALUES ($1, $2, $3, $4)
	    				ON CONFLICT DO NOTHING;`,
		article.SourceID,
		article.Title,
		article.Link,
		article.PublishedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *ArticlePostgresStorage) AllNotPosted(ctx context.Context, since time.Time, limit uint64) ([]model.Article, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `SELECT 
				a.id AS a_id, 
				s.id AS s_id,
				a.title AS a_title,
				a.link AS a_link,
				a.published_at AS a_published_at,
				a.posted_at AS a_posted_at,
				a.created_at AS a_created_at
			FROM articles a JOIN sources s ON s.id = a.source_id
			WHERE a.posted_at IS NULL 
				AND a.published_at >= $1::timestamp
			ORDER BY a.created_at DESC, s.priority DESC LIMIT $2;`

	rows, err := conn.QueryxContext(ctx, query, since.UTC().Format(time.RFC3339), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article

	for rows.Next() {
		var a model.Article
		var postedAt sql.NullTime

		if err := rows.Scan(
			&a.ID,
			&a.SourceID,
			&a.Title,
			&a.Link,
			&a.PublishedAt,
			&postedAt,
			&a.CreatedAt,
		); err != nil {
			return nil, err
		}

		a.PostedAt = postedAt.Time
		articles = append(articles, a)
	}

	return articles, nil
}

func (s *ArticlePostgresStorage) MarkAsPosted(ctx context.Context, article model.Article) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`UPDATE articles SET posted_at = $1::timestamp WHERE id = $2;`,
		time.Now().UTC().Format(time.RFC3339),
		article.ID,
	); err != nil {
		return err
	}

	return nil
}
package model

import (
	"time"
)

type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	FeedURL   string    `db:"feed_url"`
	Priority  int       `db:"priority"`
	CreatedAt time.Time `db:"created_at"`
}


type Article struct {
	ID          int64
	SourceID    int64
	Title       string
	Link        string
	Summary     string
	PublishedAt time.Time
	PostedAt    time.Time
	CreatedAt   time.Time
}

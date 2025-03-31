package models

import (
	"time"
)

// Article represents a news article
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Author    string    `json:"author"`
	Points    int       `json:"points"`
	Comments  int       `json:"comments"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}

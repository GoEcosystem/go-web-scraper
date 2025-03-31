package models

import (
	"time"
)

// Product represents a product item
type Product struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	URL         string    `json:"url"`
	Category    string    `json:"category"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"created_at"`
}

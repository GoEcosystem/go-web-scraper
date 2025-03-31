package models

import (
	"time"
)

// Scraper represents a scraper configuration
type Scraper struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	LastRun     *time.Time `json:"last_run"`
	Status      string     `json:"status"`
}

// ScraperRun represents a single execution of a scraper
type ScraperRun struct {
	ID        int64     `json:"id"`
	ScraperID int64     `json:"scraper_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	Items     int       `json:"items"`
	Error     string    `json:"error"`
}

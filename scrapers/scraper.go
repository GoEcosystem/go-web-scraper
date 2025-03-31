package scrapers

import (
	"fmt"
	"strings"

	"github.com/GoEcosystem/go-web-scraper/models"
)

// ScrapeTarget represents the target website to scrape
type ScrapeTarget string

const (
	// TargetBookstore is the bookstore website target
	TargetBookstore ScrapeTarget = "bookstore"
	// TargetHackerNews is the Hacker News website target
	TargetHackerNews ScrapeTarget = "hackernews"
)

// Scrape scrapes data from the specified target
func Scrape(target ScrapeTarget, pages int) (interface{}, error) {
	switch strings.ToLower(string(target)) {
	case string(TargetBookstore):
		fmt.Printf("Starting bookstore scraper for %d pages\n", pages)
		scraper := NewBookstoreScraper()
		return scraper.Scrape(pages)
	case string(TargetHackerNews):
		fmt.Printf("Starting Hacker News scraper for %d pages\n", pages)
		scraper := NewHackerNewsScraper()
		return scraper.Scrape(pages)
	default:
		return nil, fmt.Errorf("unknown scrape target: %s", target)
	}
}

// GetScraperInfo returns information about available scrapers
func GetScraperInfo() []models.Scraper {
	return []models.Scraper{
		{
			Name:        "Bookstore",
			Description: "Scrapes product information from a demo bookstore",
			Type:        string(TargetBookstore),
			URL:         "https://books.toscrape.com",
			LastRun:     nil,
			Status:      "ready",
		},
		{
			Name:        "Hacker News",
			Description: "Scrapes the latest articles from Hacker News",
			Type:        string(TargetHackerNews),
			URL:         "https://news.ycombinator.com",
			LastRun:     nil,
			Status:      "ready",
		},
	}
}

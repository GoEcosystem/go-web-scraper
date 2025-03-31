package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

// List of user agents to rotate through
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
}

// CreateCollector creates a new colly collector with configured settings
func CreateCollector(domain string) *colly.Collector {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a collector
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent(getRandomUserAgent()),
	)

	// Configure limits for ethical scraping
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	return c
}

// getRandomUserAgent returns a random user agent from the list
func getRandomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

// ParseFloat parses a string to a float64, returning 0 if parsing fails
func ParseFloat(s string) (float64, error) {
	// Use strconv.ParseFloat to parse the string
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

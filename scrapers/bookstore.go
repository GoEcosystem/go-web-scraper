package scrapers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"

	"github.com/GoEcosystem/go-web-scraper/models"
	"github.com/GoEcosystem/go-web-scraper/utils"
)

const (
	bookstoreBaseURL = "https://books.toscrape.com/catalogue/"
)

// BookstoreScraper is a scraper for the bookstore website
type BookstoreScraper struct {
	collector *colly.Collector
	products  []models.Product
}

// NewBookstoreScraper creates a new bookstore scraper
func NewBookstoreScraper() *BookstoreScraper {
	return &BookstoreScraper{
		collector: utils.CreateCollector("books.toscrape.com"),
		products:  make([]models.Product, 0),
	}
}

// Scrape scrapes products from the bookstore
func (s *BookstoreScraper) Scrape(pages int) ([]models.Product, error) {
	if pages <= 0 {
		pages = 1
	}

	// Setup callbacks
	s.setupCallbacks()

	// Start scraping from the first page
	startURL := "https://books.toscrape.com/catalogue/page-1.html"
	fmt.Printf("Starting to scrape bookstore from %s\n", startURL)
	err := s.collector.Visit(startURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start scraping: %w", err)
	}

	// Visit additional pages if requested
	for i := 2; i <= pages; i++ {
		nextPageURL := fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", i)
		fmt.Printf("Visiting page %d: %s\n", i, nextPageURL)
		err := s.collector.Visit(nextPageURL)
		if err != nil {
			fmt.Printf("Error visiting page %d: %s\n", i, err)
			break
		}
	}

	// Wait for scraping to finish
	s.collector.Wait()

	fmt.Printf("Scraped %d products from the bookstore\n", len(s.products))
	return s.products, nil
}

// setupCallbacks configures the collector callbacks
func (s *BookstoreScraper) setupCallbacks() {
	// Process each product
	s.collector.OnHTML("article.product_pod", func(e *colly.HTMLElement) {
		// Extract product details
		title := e.ChildAttr("h3 a", "title")
		url := e.ChildAttr("h3 a", "href")
		price := e.ChildText("div.product_price p.price_color")
		availability := e.ChildText("div.product_price p.instock.availability")
		ratingClass := e.ChildAttr("p.star-rating", "class")
		imageURL := e.ChildAttr("div.image_container img", "src")
		
		// Make URL absolute if it's relative
		if !strings.HasPrefix(url, "http") {
			url = bookstoreBaseURL + url
		}

		// Make image URL absolute if it's relative
		if strings.HasPrefix(imageURL, "../") {
			imageURL = "https://books.toscrape.com/" + strings.TrimPrefix(imageURL, "../")
		}

		// Parse price
		priceValue := 0.0
		if priceStr := strings.TrimPrefix(price, "£"); priceStr != "" {
			priceValue, _ = utils.ParseFloat(priceStr)
		}

		// Extract category
		category := "Books"
		if strings.Contains(ratingClass, "One") {
			// Use rating class as a proxy for extracting category
			category = "Fiction"
		} else if strings.Contains(ratingClass, "Three") {
			category = "Non-fiction"
		} else if strings.Contains(ratingClass, "Five") {
			category = "Bestseller"
		}
		
		// Create description from availability
		description := fmt.Sprintf("A book with %s availability.", availability)
		
		// Create product
		product := models.Product{
			Title:       title,
			Description: description,
			Price:       priceValue,
			ImageURL:    imageURL,
			URL:         url,
			Category:    category,
			Source:      "books.toscrape.com",
			CreatedAt:   time.Now(),
		}

		// Add to products if we have a title
		if title != "" {
			s.products = append(s.products, product)
		}
	})

	// Debug: Print the URL of each page visited
	s.collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	// Debug: Print errors
	s.collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error visiting %s: %s\n", r.Request.URL, err)
	})
}

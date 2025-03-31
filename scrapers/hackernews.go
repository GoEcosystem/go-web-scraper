package scrapers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/GoEcosystem/go-web-scraper/models"
	"github.com/GoEcosystem/go-web-scraper/utils"
)

// HackerNewsScraper scrapes articles from Hacker News
type HackerNewsScraper struct {
	collector *colly.Collector
	articles  []models.Article
}

// NewHackerNewsScraper creates a new Hacker News scraper
func NewHackerNewsScraper() *HackerNewsScraper {
	return &HackerNewsScraper{
		collector: utils.CreateCollector("news.ycombinator.com"),
		articles:  make([]models.Article, 0),
	}
}

// Scrape scrapes the top stories from Hacker News
func (s *HackerNewsScraper) Scrape(pages int) ([]models.Article, error) {
	if pages <= 0 {
		pages = 1
	}

	// Setup callbacks
	s.setupCallbacks()

	// Start scraping from the first page
	startURL := "https://news.ycombinator.com/news"
	fmt.Printf("Starting to scrape Hacker News from %s\n", startURL)
	err := s.collector.Visit(startURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start scraping: %w", err)
	}

	// Visit additional pages if requested
	for i := 2; i <= pages; i++ {
		nextPageURL := fmt.Sprintf("https://news.ycombinator.com/news?p=%d", i)
		fmt.Printf("Visiting page %d: %s\n", i, nextPageURL)
		err := s.collector.Visit(nextPageURL)
		if err != nil {
			fmt.Printf("Error visiting page %d: %s\n", i, err)
			break
		}
	}

	// Wait for scraping to finish
	s.collector.Wait()

	fmt.Printf("Scraped %d articles from Hacker News\n", len(s.articles))
	return s.articles, nil
}

// setupCallbacks configures the collector callbacks
func (s *HackerNewsScraper) setupCallbacks() {
	// Create a map to temporarily store article data
	articleMap := make(map[string]*models.Article)

	// Process each story row (class "athing")
	s.collector.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		if id == "" {
			return
		}

		title := e.ChildText(".titleline > a")
		url := e.ChildAttr(".titleline > a", "href")
		
		// If URL is relative, make it absolute
		if strings.HasPrefix(url, "item?id=") {
			url = fmt.Sprintf("https://news.ycombinator.com/%s", url)
		}

		// Find the site domain
		source := e.ChildText(".sitestr")
		if source == "" {
			source = "news.ycombinator.com"
		}

		// Create a new article and store it in the map
		articleMap[id] = &models.Article{
			Title:     title,
			URL:       url,
			Source:    source,
			CreatedAt: time.Now(),
		}
		
		fmt.Printf("Found article: %s (ID: %s)\n", title, id)
	})

	// Process the subtext rows that come after the story rows
	s.collector.OnHTML("tr:not(.athing)", func(e *colly.HTMLElement) {
		// Check if this row has a score element, which indicates it's a subtext row
		scoreElem := e.DOM.Find(".score")
		if scoreElem.Length() == 0 {
			return
		}

		// Extract the ID from the previous row
		prevElem := e.DOM.Prev()
		id, _ := prevElem.Attr("id")
		if id == "" {
			return
		}

		// Check if we have this article in our map
		article, exists := articleMap[id]
		if !exists {
			return
		}

		// Extract points
		pointsStr := e.ChildText(".score")
		if pointsStr != "" {
			pointsStr = strings.TrimSuffix(pointsStr, " points")
			pointsStr = strings.TrimSuffix(pointsStr, " point")
			points, _ := strconv.Atoi(pointsStr)
			article.Points = points
		}

		// Extract author
		article.Author = e.ChildText(".hnuser")

		// Extract comments count from the last link
		commentsText := ""
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Text, "comment") {
				commentsText = el.Text
			}
		})

		if commentsText != "" {
			commentsText = strings.TrimSuffix(commentsText, " comments")
			commentsText = strings.TrimSuffix(commentsText, " comment")
			comments, _ := strconv.Atoi(commentsText)
			article.Comments = comments
		}

		// Add the article to our final collection
		s.articles = append(s.articles, *article)
		
		fmt.Printf("Added article: %s by %s with %d points\n", article.Title, article.Author, article.Points)
		
		// Remove from map to free memory
		delete(articleMap, id)
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

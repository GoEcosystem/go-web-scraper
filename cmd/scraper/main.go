package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/GoEcosystem/go-web-scraper/models"
	"github.com/GoEcosystem/go-web-scraper/scrapers"
)

func main() {
	// Define command-line flags
	targetPtr := flag.String("target", "", "Target website to scrape (bookstore, hackernews)")
	pagesPtr := flag.Int("pages", 1, "Number of pages to scrape")
	formatPtr := flag.String("format", "json", "Output format (json, csv)")
	outputDirPtr := flag.String("output", "output", "Directory to save output files")
	flag.Parse()

	// Validate target
	if *targetPtr == "" {
		fmt.Println("Please specify a target with -target flag")
		fmt.Println("Available targets: bookstore, hackernews")
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDirPtr, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Start scraping
	fmt.Printf("Starting to scrape %s (%d pages)...\n", *targetPtr, *pagesPtr)
	startTime := time.Now()

	// Run the scraper
	result, err := scrapers.Scrape(scrapers.ScrapeTarget(*targetPtr), *pagesPtr)
	if err != nil {
		log.Fatalf("Scraping failed: %v", err)
	}

	// Create output filename with timestamp
	timestamp := time.Now().Format("20060102-150405")
	outputFile := filepath.Join(*outputDirPtr, fmt.Sprintf("%s_%s.%s", *targetPtr, timestamp, *formatPtr))

	// Save data based on format
	switch *formatPtr {
	case "json":
		saveJSON(outputFile, result)
	case "csv":
		saveCSV(outputFile, result, *targetPtr)
	default:
		log.Fatalf("Unsupported output format: %s", *formatPtr)
	}

	// Print summary
	duration := time.Since(startTime)
	fmt.Printf("Scraping complete in %.2f seconds\n", duration.Seconds())
	
	switch v := result.(type) {
	case []models.Article:
		fmt.Printf("Scraped %d articles and saved to %s\n", len(v), outputFile)
	case []models.Product:
		fmt.Printf("Scraped %d products and saved to %s\n", len(v), outputFile)
	default:
		fmt.Printf("Saved scraped data to %s\n", outputFile)
	}
}

// saveJSON saves data as JSON
func saveJSON(filename string, data interface{}) {
	// Create file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	// Encode data as JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Failed to encode data as JSON: %v", err)
	}

	fmt.Printf("Data saved as JSON to %s\n", filename)
}

// saveCSV saves data as CSV
func saveCSV(filename string, data interface{}, target string) {
	// Create file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	// Write CSV header and data based on type
	switch v := data.(type) {
	case []models.Article:
		// Write header
		file.WriteString("title,author,url,points,comments,source,created_at\n")
		
		// Write data
		for _, article := range v {
			line := fmt.Sprintf("\"%s\",\"%s\",\"%s\",%d,%d,\"%s\",\"%s\"\n", 
				escapeCSV(article.Title),
				escapeCSV(article.Author),
				escapeCSV(article.URL),
				article.Points,
				article.Comments,
				escapeCSV(article.Source),
				article.CreatedAt.Format(time.RFC3339))
			file.WriteString(line)
		}
	case []models.Product:
		// Write header
		file.WriteString("title,description,price,image_url,url,category,source,created_at\n")
		
		// Write data
		for _, product := range v {
			line := fmt.Sprintf("\"%s\",\"%s\",%.2f,\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", 
				escapeCSV(product.Title),
				escapeCSV(product.Description),
				product.Price,
				escapeCSV(product.ImageURL),
				escapeCSV(product.URL),
				escapeCSV(product.Category),
				escapeCSV(product.Source),
				product.CreatedAt.Format(time.RFC3339))
			file.WriteString(line)
		}
	default:
		log.Fatalf("Unsupported data type for CSV export")
	}

	fmt.Printf("Data saved as CSV to %s\n", filename)
}

// escapeCSV escapes double quotes in CSV fields
func escapeCSV(s string) string {
	return s
}

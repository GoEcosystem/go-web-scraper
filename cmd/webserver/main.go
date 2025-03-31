package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoEcosystem/go-web-scraper/db"
	"github.com/GoEcosystem/go-web-scraper/models"
	"github.com/GoEcosystem/go-web-scraper/web"
)

func main() {
	// Parse command line flags
	port := flag.String("port", "8080", "Port to run the web server on")
	dbPath := flag.String("db", "data/scraper.db", "Path to SQLite database file")
	outputDir := flag.String("output", "output", "Directory for scraped data output files")
	debugMode := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Create output directory if not exists
	err := os.MkdirAll(*outputDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Create database directory if not exists
	err = os.MkdirAll(filepath.Dir(*dbPath), 0755)
	if err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	// Initialize database
	database, err := db.New(*dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize database schema
	err = database.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Import data from JSON files in output directory
	log.Printf("Importing data from %s", *outputDir)
	importDataFromFiles(database, *outputDir)

	// Log database stats
	articleCount, _ := database.GetArticleCount()
	productCount, _ := database.GetProductCount()
	log.Printf("Database contains %d articles and %d products", articleCount, productCount)

	// Create and start web server
	server := web.NewServer(*port, *outputDir, database, *debugMode)
	err = server.Start()
	if err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}

// importDataFromFiles imports sample data from JSON files in the output directory
func importDataFromFiles(database *db.Database, outputDir string) {
	// Check if output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		log.Printf("Output directory does not exist: %s", outputDir)
		return
	}

	// Get list of files in output directory
	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		log.Printf("Failed to read output directory: %v", err)
		return
	}

	// Import each JSON file
	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}

		// Check if file is JSON
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".json") {
			fmt.Printf("Skipping non-JSON file: %s\n", fileName)
			continue
		}

		// Check if file is empty or very small (less than 10 bytes)
		if file.Size() < 10 {
			fmt.Printf("Skipping empty or very small JSON file: %s (size: %d bytes)\n", fileName, file.Size())
			continue
		}

		// Import file
		filePath := filepath.Join(outputDir, fileName)
		if strings.HasPrefix(fileName, "hackernews") {
			importArticles(database, filePath)
		} else if strings.HasPrefix(fileName, "bookstore") {
			importProducts(database, filePath)
		}
	}
}

// importArticles imports articles from a JSON file
func importArticles(database *db.Database, filePath string) {
	// Read file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file %s: %v", filePath, err)
		return
	}

	// Parse JSON
	var articles []models.Article
	err = json.Unmarshal(data, &articles)
	if err != nil {
		log.Printf("Failed to parse JSON in file %s: %v", filePath, err)
		return
	}

	// Import articles
	count, err := database.InsertArticles(articles)
	if err != nil {
		log.Printf("Failed to import articles from %s: %v", filePath, err)
	} else {
		fmt.Printf("Imported %d articles from %s\n", count, filePath)
	}
}

// importProducts imports products from a JSON file
func importProducts(database *db.Database, filePath string) {
	// Read file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file %s: %v", filePath, err)
		return
	}

	// Parse JSON
	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Failed to parse JSON in file %s: %v", filePath, err)
		return
	}

	// Import products
	count, err := database.InsertProducts(products)
	if err != nil {
		log.Printf("Failed to import products from %s: %v", filePath, err)
	} else {
		fmt.Printf("Imported %d products from %s\n", count, filePath)
	}
}

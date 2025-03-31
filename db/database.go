package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/GoEcosystem/go-web-scraper/models"
	_ "github.com/mattn/go-sqlite3"
)

// Database represents a connection to the SQLite database
type Database struct {
	db         *sql.DB
	dbPath     string
	initialized bool
}

// New creates a new database connection
func New(dbPath string) (*Database, error) {
	// Create database directory if not exists
	err := os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create database directory: %v", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Set connection parameters
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	// Create instance
	database := &Database{
		db:     db,
		dbPath: dbPath,
	}

	// Check connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return database, nil
}

// Initialize creates database tables if they don't exist
func (d *Database) Initialize() error {
	if d.initialized {
		return nil
	}

	// Create articles table
	_, err := d.db.Exec(`
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		url TEXT,
		author TEXT,
		points INTEGER,
		comments INTEGER,
		source TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return fmt.Errorf("failed to create articles table: %v", err)
	}

	// Create products table
	_, err = d.db.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		price REAL,
		image_url TEXT,
		url TEXT,
		category TEXT,
		source TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return fmt.Errorf("failed to create products table: %v", err)
	}

	// Create scrapers table
	_, err = d.db.Exec(`
	CREATE TABLE IF NOT EXISTS scrapers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT NOT NULL,
		description TEXT,
		url TEXT,
		last_run TIMESTAMP,
		status TEXT
	)
	`)
	if err != nil {
		return fmt.Errorf("failed to create scrapers table: %v", err)
	}

	// Create scraper_runs table
	_, err = d.db.Exec(`
	CREATE TABLE IF NOT EXISTS scraper_runs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		scraper_id INTEGER,
		start_time TIMESTAMP,
		end_time TIMESTAMP,
		status TEXT,
		items INTEGER,
		error TEXT,
		FOREIGN KEY (scraper_id) REFERENCES scrapers (id)
	)
	`)
	if err != nil {
		return fmt.Errorf("failed to create scraper_runs table: %v", err)
	}

	// Insert default scrapers if none exist
	var count int
	err = d.db.QueryRow("SELECT COUNT(*) FROM scrapers").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to count scrapers: %v", err)
	}

	if count == 0 {
		_, err = d.db.Exec(`
		INSERT INTO scrapers (name, type, description, url, status)
		VALUES 
		('Hacker News', 'hackernews', 'Scrapes the latest articles from Hacker News', 'https://news.ycombinator.com', 'ready'),
		('Bookstore', 'bookstore', 'Scrapes product information from a demo bookstore', 'https://books.toscrape.com', 'ready')
		`)
		if err != nil {
			return fmt.Errorf("failed to insert default scrapers: %v", err)
		}
	}

	d.initialized = true
	log.Println("Database initialized successfully")
	return nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.db.Close()
}

// GetArticles retrieves a list of articles with pagination
func (d *Database) GetArticles(page, pageSize int) ([]models.Article, error) {
	offset := (page - 1) * pageSize

	// Query articles
	rows, err := d.db.Query(`
		SELECT 
			id, title, url, author, points, comments, source, created_at
		FROM articles
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %v", err)
	}
	defer rows.Close()

	// Parse results
	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.URL,
			&article.Author,
			&article.Points,
			&article.Comments,
			&article.Source,
			&article.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan article row: %v", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// GetProducts retrieves a list of products with pagination
func (d *Database) GetProducts(page, pageSize int) ([]models.Product, error) {
	offset := (page - 1) * pageSize

	// Query products
	rows, err := d.db.Query(`
		SELECT 
			id, title, description, price, image_url, url, category, source, created_at
		FROM products
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %v", err)
	}
	defer rows.Close()

	// Parse results
	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.ImageURL,
			&product.URL,
			&product.Category,
			&product.Source,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// GetArticleCount returns the total number of articles
func (d *Database) GetArticleCount() (int, error) {
	var count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM articles").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count articles: %v", err)
	}
	return count, nil
}

// GetProductCount returns the total number of products
func (d *Database) GetProductCount() (int, error) {
	var count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count products: %v", err)
	}
	return count, nil
}

// GetScrapers returns all scrapers
func (d *Database) GetScrapers() ([]models.Scraper, error) {
	// Query scrapers
	rows, err := d.db.Query(`
		SELECT 
			id, name, type, description, url, last_run, status
		FROM scrapers
		ORDER BY name
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query scrapers: %v", err)
	}
	defer rows.Close()

	// Parse results
	scrapers := []models.Scraper{}
	for rows.Next() {
		var scraper models.Scraper
		err := rows.Scan(
			&scraper.ID,
			&scraper.Name,
			&scraper.Type,
			&scraper.Description,
			&scraper.URL,
			&scraper.LastRun,
			&scraper.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan scraper row: %v", err)
		}
		scrapers = append(scrapers, scraper)
	}

	return scrapers, nil
}

// InsertArticles inserts multiple articles into the database
func (d *Database) InsertArticles(articles []models.Article) (int, error) {
	// Begin transaction
	tx, err := d.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare statement
	stmt, err := tx.Prepare(`
		INSERT INTO articles (title, url, author, points, comments, source, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Insert articles
	count := 0
	for _, article := range articles {
		// Set created_at to current time if not provided
		createdAt := article.CreatedAt
		if createdAt.IsZero() {
			createdAt = time.Now()
		}

		_, err := stmt.Exec(
			article.Title,
			article.URL,
			article.Author,
			article.Points,
			article.Comments,
			article.Source,
			createdAt,
		)
		if err != nil {
			return count, fmt.Errorf("failed to insert article: %v", err)
		}
		count++
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return count, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return count, nil
}

// InsertProducts inserts multiple products into the database
func (d *Database) InsertProducts(products []models.Product) (int, error) {
	// Begin transaction
	tx, err := d.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare statement
	stmt, err := tx.Prepare(`
		INSERT INTO products (title, description, price, image_url, url, category, source, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Insert products
	count := 0
	for _, product := range products {
		// Set created_at to current time if not provided
		createdAt := product.CreatedAt
		if createdAt.IsZero() {
			createdAt = time.Now()
		}

		_, err := stmt.Exec(
			product.Title,
			product.Description,
			product.Price,
			product.ImageURL,
			product.URL,
			product.Category,
			product.Source,
			createdAt,
		)
		if err != nil {
			return count, fmt.Errorf("failed to insert product: %v", err)
		}
		count++
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return count, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return count, nil
}

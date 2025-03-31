package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/GoEcosystem/go-web-scraper/db"
	"github.com/GoEcosystem/go-web-scraper/models"
)

// Server represents a web server
type Server struct {
	Port      string
	OutputDir string
	db        *db.Database
	templates *template.Template
	Debug     bool
	PageSize  int
}

// PageData holds data for template rendering
type PageData struct {
	Title        string
	Articles     []models.Article
	Products     []models.Product
	Files        []FileInfo
	Error        string
	Debug        string
	HasPrevPage  bool
	HasNextPage  bool
	ArticleCount int
	ProductCount int
	Scrapers     []models.Scraper
	Message      string
	Command      string
	CurrentPage  int
	TotalPages   int
}

// FileInfo holds information about a scraped data file
type FileInfo struct {
	Name     string
	Path     string
	Type     string
	Modified time.Time
	Size     int64
}

// NewServer creates a new web server
func NewServer(port string, outputDir string, db *db.Database, debug bool) *Server {
	// Create server instance
	server := &Server{
		Port:      port,
		OutputDir: outputDir,
		db:        db,
		Debug:     debug,
		PageSize:  10,
	}

	// Define template functions
	funcMap := template.FuncMap{
		"formatTime": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"formatPrice": func(p float64) string {
			return fmt.Sprintf("$%.2f", p)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"seq": func(start, end int) []int {
			var result []int
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
			return result
		},
	}

	// Find template files
	templateDir := filepath.Join("web", "templates")
	templateFiles, err := filepath.Glob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		log.Fatalf("Error finding templates: %v", err)
	}
	log.Printf("Found %d template files: %v", len(templateFiles), templateFiles)

	// Parse all templates together
	templates, err := template.New("").Funcs(funcMap).ParseFiles(templateFiles...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	server.templates = templates

	// Log loaded templates for debugging
	log.Printf("Successfully loaded %d templates", len(templates.Templates()))
	for _, t := range templates.Templates() {
		log.Printf("Loaded template: %s", t.Name())
	}

	return server
}

// Start starts the web server
func (s *Server) Start() error {
	// Set up routes
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/articles", s.handleArticles)
	http.HandleFunc("/products", s.handleProducts)
	http.HandleFunc("/static/", s.handleStatic)
	
	// API endpoints
	http.HandleFunc("/api/scrapers/run", s.handleRunScraper)
	http.HandleFunc("/api/scrapers/status", s.handleScraperStatus)

	// Start server
	log.Printf("Starting web server on http://localhost:%s\n", s.Port)
	fmt.Printf("Starting web server on http://localhost:%s\n", s.Port)
	return http.ListenAndServe(":"+s.Port, nil)
}

// handleIndex handles the index page request
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	// Redirect /index to /
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get database statistics
	articleCount, err := s.db.GetArticleCount()
	if err != nil {
		s.renderError(w, "Error getting article count", err)
		return
	}

	productCount, err := s.db.GetProductCount()
	if err != nil {
		s.renderError(w, "Error getting product count", err)
		return
	}

	// Get scrapers
	scrapers, err := s.db.GetScrapers()
	if err != nil {
		s.renderError(w, "Error getting scrapers", err)
		return
	}

	// Get output files
	files, err := s.getOutputFiles()
	if err != nil {
		log.Printf("Error getting output files: %v", err)
		// Continue with empty files slice
		files = []FileInfo{}
	}

	// Log details for debugging
	log.Printf("Dashboard data: Files=%d, Articles=%d, Products=%d, Scrapers=%d", 
		len(files), articleCount, productCount, len(scrapers))

	// Prepare data for template
	data := PageData{
		Title:        "Web Scraper Dashboard",
		Files:        files,
		ArticleCount: articleCount,
		ProductCount: productCount,
		Scrapers:     scrapers,
		// Set Articles and Products to nil to avoid triggering empty state messages
		Articles:     nil,
		Products:     nil,
		// Set a custom message for the dashboard if needed
		Message:      "",
		Command:      "",
	}

	// Render template directly using index.html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.Must(s.templates.Clone()).ParseFiles("web/templates/index.html", "web/templates/layout.html"))
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("Error executing index template: %v", err)
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
	} else {
		log.Printf("Successfully rendered index template")
	}
}

// handleArticles handles the articles page request
func (s *Server) handleArticles(w http.ResponseWriter, r *http.Request) {
	// Get page number from query params
	page := 1
	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		pageNum, err := strconv.Atoi(pageParam)
		if err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	// Set page size
	pageSize := 10

	// Get articles from database
	articles, err := s.db.GetArticles(page, pageSize)
	if err != nil {
		s.renderError(w, "Error getting articles", err)
		return
	}

	// Log article count for debugging
	log.Printf("Retrieved %d articles for page %d", len(articles), page)

	// Print detailed info about the first few articles for debugging
	if len(articles) > 0 {
		log.Printf("Articles are available - should display properly")
		for i, article := range articles {
			if i < 3 { // Only print details for the first 3 articles to avoid flooding logs
				log.Printf("Article %d: Title=%s, Author=%s, Points=%d, Comments=%d", 
					i+1, article.Title, article.Author, article.Points, article.Comments)
			}
		}
	} else {
		// Get article count for warning message
		articleCount, err := s.db.GetArticleCount()
		if err != nil {
			log.Printf("WARNING: Failed to get article count: %v", err)
		} else {
			log.Printf("WARNING: No articles returned from database despite having %d total articles", articleCount)
		}
	}

	// Get total count for pagination
	totalArticles, err := s.db.GetArticleCount()
	if err != nil {
		s.renderError(w, "Error getting article count", err)
		return
	}

	// Calculate total pages
	totalPages := (totalArticles + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	// Custom rendering to handle the empty state properly
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Prepare data for template
	data := PageData{
		Title:        "Articles",
		Articles:     articles,
		CurrentPage:  page,
		TotalPages:   totalPages,
		HasPrevPage:  page > 1,
		HasNextPage:  page < totalPages,
		ArticleCount: totalArticles,
		// Ensure no products or product-related messages
		Products:     nil,        // Must be nil to avoid triggering product section
		ProductCount: 0,          // Set to 0 to avoid product-related messages
		Message:      "No articles found. Please run the scraper to collect data. You can do this by running the following command in your terminal:\ngo run cmd/scraper/main.go -target hackernews -pages 3 -format json",
		Command:      "go run cmd/scraper/main.go -target hackernews -pages 3 -format json",
	}

	// Execute the template for articles
	tmpl := template.Must(template.Must(s.templates.Clone()).ParseFiles("web/templates/articles.html", "web/templates/layout.html"))
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("Error executing articles.html template: %v", err)
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
	} else {
		log.Printf("Successfully rendered articles template with %d articles", len(articles))
	}
}

// handleProducts handles the products page request
func (s *Server) handleProducts(w http.ResponseWriter, r *http.Request) {
	// Get page number from query params
	page := 1
	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		pageNum, err := strconv.Atoi(pageParam)
		if err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	// Set page size
	pageSize := 10

	// Get products from database
	products, err := s.db.GetProducts(page, pageSize)
	if err != nil {
		s.renderError(w, "Error getting products", err)
		return
	}

	// Log product count for debugging
	log.Printf("Retrieved %d products for page %d", len(products), page)

	// Get total count for pagination
	totalProducts, err := s.db.GetProductCount()
	if err != nil {
		s.renderError(w, "Error getting product count", err)
		return
	}

	// Calculate total pages
	totalPages := (totalProducts + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	// Prepare data for template
	data := PageData{
		Title:        "Products",
		Products:     products,
		CurrentPage:  page,
		TotalPages:   totalPages,
		HasPrevPage:  page > 1,
		HasNextPage:  page < totalPages,
		ProductCount: totalProducts,
		// Ensure no article-related messages
		Articles:     nil,
		ArticleCount: 0,
		// Set specific message for products
		Message:      "No products found. Please run the scraper to collect data. You can do this by running the following command in your terminal:",
		Command:      "go run cmd/scraper/main.go -target bookstore -pages 3 -format json",
	}

	// Execute template
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.Must(s.templates.Clone()).ParseFiles("web/templates/products.html", "web/templates/layout.html"))
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("Error executing products.html template: %v", err)
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
	} else {
		log.Printf("Successfully rendered products template with %d products", len(products))
	}
}

// handleStatic handles static file requests
func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
	// Get file path
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join("web", "static", filePath)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Serve static file
	http.ServeFile(w, r, fullPath)
}

// handleRunScraper handles the API request to run a scraper
func (s *Server) handleRunScraper(w http.ResponseWriter, r *http.Request) {
	// Check method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get scraper ID
	scraperID := r.FormValue("scraper_id")
	if scraperID == "" {
		http.Error(w, "Scraper ID is required", http.StatusBadRequest)
		return
	}

	// Simulate running a scraper (in a real app, this would start a goroutine)
	runID := fmt.Sprintf("%d", time.Now().UnixNano())
	
	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"run_id": runID,
	})
}

// handleScraperStatus handles the API request to check scraper status
func (s *Server) handleScraperStatus(w http.ResponseWriter, r *http.Request) {
	// Get run ID
	runID := r.URL.Query().Get("run_id")
	if runID == "" {
		http.Error(w, "Run ID is required", http.StatusBadRequest)
		return
	}

	// Simulate checking status (in a real app, this would check a database)
	status := "success"
	
	// Return status response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": status,
	})
}

// renderError renders an error page
func (s *Server) renderError(w http.ResponseWriter, errorMsg string, err error) {
	log.Printf("Error: %s: %v", errorMsg, err)
	
	data := PageData{
		Title: "Error",
		Error: errorMsg,
		Message: "", // Ensure no default message
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	tmpl := template.Must(template.Must(s.templates.Clone()).ParseFiles("web/templates/error.html", "web/templates/layout.html"))
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error rendering error template: %v", err)
		http.Error(w, fmt.Sprintf("%s: %v", errorMsg, err), http.StatusInternalServerError)
	}
}

// getOutputFiles gets a list of JSON files in the output directory
func (s *Server) getOutputFiles() ([]FileInfo, error) {
	// Check if output directory exists
	if _, err := os.Stat(s.OutputDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("output directory does not exist: %s", s.OutputDir)
	}

	// Get list of files in output directory
	entries, err := os.ReadDir(s.OutputDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read output directory: %v", err)
	}

	// Filter JSON files
	files := []FileInfo{}
	fmt.Printf("Looking for files in: %s\n", s.OutputDir)
	fmt.Printf("Found %d files/directories in output directory\n", len(entries))
	
	for _, entry := range entries {
		// Skip directories
		if entry.IsDir() {
			continue
		}

		// Check if file is JSON
		fileName := entry.Name()
		if !strings.HasSuffix(fileName, ".json") {
			fmt.Printf("Skipping non-JSON file: %s\n", fileName)
			continue
		}

		// Check if file is empty or very small (less than 10 bytes)
		path := filepath.Join(s.OutputDir, fileName)
		info, err := os.Stat(path)
		if err != nil {
			continue
		}
		if info.Size() < 10 {
			fmt.Printf("Skipping empty or very small JSON file: %s (size: %d bytes)\n", fileName, info.Size())
			continue
		}

		// Determine file type (articles or products)
		fileType := "unknown"
		if strings.HasPrefix(fileName, "hackernews") {
			fileType = "articles"
		} else if strings.HasPrefix(fileName, "bookstore") {
			fileType = "products"
		}

		// Create file info
		fmt.Printf("Added file to list: %s (%s)\n", fileName, fileType)
		files = append(files, FileInfo{
			Name:     fileName,
			Path:     path,
			Type:     fileType,
			Modified: info.ModTime(),
			Size:     info.Size(),
		})
	}

	return files, nil
}

---
layout: default
title: Go Web Scraper
---

# Go Web Scraper

A comprehensive web scraping solution built in Go, featuring concurrent scraping, database integration, and a responsive web interface.

## Features

- **Multi-site Scraping**: Extract data from different websites including Hacker News and bookstore sites
- **Concurrent Processing**: Utilizes Go's goroutines for efficient parallel scraping
- **SQLite Database**: Persistent storage of all scraped data with proper schema design
- **Web Interface**: Bootstrap-styled UI for browsing and searching scraped content
- **Data Export**: Export scraped data to JSON and CSV formats
- **Pagination Support**: Efficient navigation through large datasets of articles and products
- **Dashboard**: Statistics and overview of the database contents
- **Ethical Scraping**: Built-in rate limiting and user agent rotation
- **Import Functionality**: Load data from JSON files into the database

## Getting Started

### Prerequisites

- Go 1.18 or higher
- SQLite

### Installation

```bash
# Clone the repository
git clone https://github.com/GoEcosystem/go-web-scraper.git
cd go-web-scraper

# Build the application
go build -o scraper ./cmd/scraper
```

### Usage

#### Command Line Scraping

```bash
# Scrape Hacker News
./scraper -source hackernews -limit 20

# Scrape Bookstore
./scraper -source bookstore -pages 3
```

#### Starting the Web Interface

```bash
# Start the web server on port 8080
./scraper -server -port 8080
```

Then open your browser and navigate to `http://localhost:8080`

## Project Structure

```
go-web-scraper/
├── cmd/                # Command-line applications
│   └── scraper/        # Main scraper application
├── models/             # Data models
├── scrapers/           # Website-specific scraper implementations
├── db/                 # Database interaction layer
├── web/                # Web interface
│   ├── templates/      # HTML templates
│   ├── static/         # CSS, JavaScript, images
│   └── handlers/       # HTTP handlers
├── utils/              # Utility functions
└── docs/               # Documentation
```

## Screenshots

*Coming soon*

## API Documentation

The web interface provides a simple API for accessing scraped data:

- `GET /api/articles` - List all articles
- `GET /api/articles/:id` - Get specific article
- `GET /api/products` - List all products
- `GET /api/products/:id` - Get specific product

For full API documentation, see the [API documentation page](https://goecosystem.github.io/go-web-scraper/api-documentation).

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

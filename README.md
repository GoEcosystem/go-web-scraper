# GoEcosystem Web Scraper

[![Go Report Card](https://goreportcard.com/badge/github.com/GoEcosystem/go-web-scraper)](https://goreportcard.com/report/github.com/GoEcosystem/go-web-scraper)
[![License](https://img.shields.io/github/license/GoEcosystem/go-web-scraper.svg)](https://github.com/GoEcosystem/go-web-scraper/blob/main/LICENSE)
[![Tests](https://github.com/GoEcosystem/go-web-scraper/workflows/Go%20Tests/badge.svg)](https://github.com/GoEcosystem/go-web-scraper/actions?query=workflow%3A%22Go+Tests%22)

A comprehensive web scraper application in Go with data persistence and web interface.

## Features

- Multi-source scraping (Hacker News, Bookstore, and more)
- Data persistence using SQLite
- Web interface for browsing scraped data
- Configurable scraping options
- Export to JSON and CSV formats
- Ethical scraping with rate limiting and user agent rotation
- Detailed logging

## Installation

### Prerequisites

- Go 1.20 or higher
- SQLite 3

### Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/GoEcosystem/go-web-scraper.git
   cd go-web-scraper
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o scraper ./cmd/scraper
   go build -o server ./cmd/webserver
   ```

## Usage

### Command Line Scraper

Run the scraper to collect data:

```bash
./scraper -target=hackernews -pages=5
./scraper -target=bookstore -pages=3
```

Available options:
- `-target`: The website to scrape (hackernews, bookstore)
- `-pages`: Number of pages to scrape
- `-output`: Output format (json, csv, db)
- `-file`: Output filename (when using json or csv)

### Web Interface

Start the web server:

```bash
./server -port=8080
```

Then open your browser at http://localhost:8080

## Project Structure

```
.
├── cmd/                  # Command-line applications
│   ├── scraper/          # CLI scraper tool
│   └── webserver/        # Web interface server
├── db/                   # Database management
├── models/               # Data models
├── scrapers/             # Website-specific scrapers
├── utils/                # Utility functions
└── web/                  # Web interface
    ├── server.go         # HTTP server
    ├── static/           # Static assets (CSS, JS)
    └── templates/        # HTML templates
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](https://github.com/GoEcosystem/go-docs/blob/main/CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

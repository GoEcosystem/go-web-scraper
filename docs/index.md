---
layout: default
title: Go Web Scraper Documentation
description: Documentation for the comprehensive web scraping solution built in Go
repo_name: go-web-scraper
---

<div class="documentation-home">
  <div class="hero-section">
    <h1>Go Web Scraper</h1>
    <p class="hero-tagline">A comprehensive web scraping solution built in Go with database integration and web interface</p>
    <div class="hero-buttons">
      <a href="https://github.com/GoEcosystem/go-web-scraper" class="button primary">View on GitHub</a>
      <a href="/go-web-scraper/api" class="button secondary">API Reference</a>
    </div>
  </div>

  <div class="features-section">
    <h2>Features</h2>
    <div class="feature-grid">
      <div class="feature-card">
        <h3>Concurrent Scraping</h3>
        <p>Utilizes Go's concurrency features for efficient data extraction from multiple sources simultaneously</p>
      </div>
      <div class="feature-card">
        <h3>SQLite Integration</h3>
        <p>Persistent storage of scraped data with a robust SQLite database backend</p>
      </div>
      <div class="feature-card">
        <h3>Web Interface</h3>
        <p>Browse and search scraped data through a clean, responsive web UI</p>
      </div>
      <div class="feature-card">
        <h3>Multiple Export Formats</h3>
        <p>Export data to JSON or CSV formats for further analysis</p>
      </div>
      <div class="feature-card">
        <h3>Rate Limiting</h3>
        <p>Built-in rate limiting to ensure ethical scraping practices</p>
      </div>
      <div class="feature-card">
        <h3>Extensible Design</h3>
        <p>Easily add new scraper modules for additional data sources</p>
      </div>
    </div>
  </div>

  <div class="documentation-section">
    <h2>Documentation</h2>
    <div class="documentation-grid">
      <a href="/go-web-scraper/api" class="doc-card">
        <h3>API Reference</h3>
        <p>Comprehensive API documentation for integrating with the Web Scraper</p>
      </a>
      <a href="/go-web-scraper/architecture" class="doc-card">
        <h3>Architecture</h3>
        <p>Understand the internal structure and design of the application</p>
      </a>
      <a href="https://github.com/GoEcosystem/go-web-scraper#installation" class="doc-card">
        <h3>Installation Guide</h3>
        <p>Instructions for setting up and running the web scraper</p>
      </a>
      <a href="https://github.com/GoEcosystem/go-web-scraper#usage" class="doc-card">
        <h3>Usage Examples</h3>
        <p>Examples showing how to use the scraper for different scenarios</p>
      </a>
    </div>
  </div>

  <div class="getting-started-section">
    <h2>Quick Start</h2>
    <div class="code-block">
      <pre><code>
# Clone the repository
git clone https://github.com/GoEcosystem/go-web-scraper.git

# Change to the project directory
cd go-web-scraper

# Build the application
go build -o scraper ./cmd/scraper

# Run a scraper
./scraper --source hackernews --output data/articles.json

# Start the web server
go run ./cmd/webserver
      </code></pre>
    </div>
  </div>

  <div class="ecosystem-section">
    <h2>GoEcosystem Projects</h2>
    <p>This project is part of the <a href="https://goecosystem.github.io">GoEcosystem</a> - a collection of high-quality Go applications, libraries, and tools.</p>
    <div class="ecosystem-links">
      <a href="https://goecosystem.github.io/go-docs" class="ecosystem-link">Go Docs</a>
      <a href="https://goecosystem.github.io/go-cheatsheets" class="ecosystem-link">Go Cheatsheets</a>
      <a href="https://goecosystem.github.io/task-api" class="ecosystem-link">Task API</a>
      <a href="https://goecosystem.github.io/go-utils" class="ecosystem-link">Go Utils</a>
    </div>
  </div>
</div>

<style>
  .documentation-home {
    max-width: 1200px;
    margin: 0 auto;
  }

  .hero-section {
    text-align: center;
    margin: 3rem 0 4rem;
  }

  .hero-tagline {
    font-size: 1.2rem;
    color: #6c757d;
    margin-bottom: 2rem;
    max-width: 800px;
    margin-left: auto;
    margin-right: auto;
  }

  .hero-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
  }

  .button {
    display: inline-block;
    padding: 0.75rem 1.5rem;
    border-radius: 4px;
    font-weight: 500;
    text-decoration: none;
    transition: all 0.2s ease-in-out;
  }

  .button.primary {
    background-color: #00ADD8;
    color: white;
  }

  .button.primary:hover {
    background-color: #0092B8;
  }

  .button.secondary {
    background-color: #f8f9fa;
    color: #212529;
    border: 1px solid #dee2e6;
  }

  .button.secondary:hover {
    background-color: #e9ecef;
  }

  .features-section, 
  .documentation-section, 
  .getting-started-section,
  .ecosystem-section {
    margin: 4rem 0;
  }

  .feature-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1.5rem;
    margin-top: 2rem;
  }

  .feature-card {
    padding: 1.5rem;
    border-radius: 5px;
    border: 1px solid #dee2e6;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }

  .feature-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0,0,0,0.05);
  }

  .feature-card h3 {
    color: #00ADD8;
    margin-top: 0;
    margin-bottom: 1rem;
  }

  .documentation-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1.5rem;
    margin-top: 2rem;
  }

  .doc-card {
    padding: 1.5rem;
    border-radius: 5px;
    border: 1px solid #dee2e6;
    text-decoration: none;
    color: inherit;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }

  .doc-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0,0,0,0.05);
    text-decoration: none;
  }

  .doc-card h3 {
    color: #00ADD8;
    margin-top: 0;
    margin-bottom: 1rem;
  }

  .code-block {
    background-color: #1F2937;
    border-radius: 5px;
    overflow: hidden;
    margin-top: 2rem;
  }

  .code-block pre {
    margin: 0;
    padding: 1.5rem;
    overflow-x: auto;
  }

  .code-block code {
    color: #E5E7EB;
    font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, Courier, monospace;
    font-size: 0.9rem;
  }

  .ecosystem-links {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    margin-top: 1.5rem;
  }

  .ecosystem-link {
    display: inline-block;
    padding: 0.5rem 1rem;
    background-color: #f8f9fa;
    border-radius: 4px;
    border: 1px solid #dee2e6;
    color: #00ADD8;
    text-decoration: none;
    transition: all 0.2s ease;
  }

  .ecosystem-link:hover {
    background-color: #e9ecef;
    text-decoration: none;
  }

  @media (max-width: 768px) {
    .hero-buttons {
      flex-direction: column;
      align-items: center;
    }

    .feature-grid,
    .documentation-grid {
      grid-template-columns: 1fr;
    }
  }
</style>

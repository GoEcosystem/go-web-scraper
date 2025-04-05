---
layout: default
title: Go Web Scraper Documentation
description: Documentation for the comprehensive web scraping solution built in Go
repo_name: go-web-scraper
---

<div class="hero">
  <img src="{{ site.baseurl }}/assets/images/logos/Go-Logo_Blue.svg" alt="Go Logo" class="hero-logo">
  <h1>Go Web Scraper</h1>
  <p>A comprehensive web scraping solution built in Go with SQLite database integration and web interface</p>
  <a href="https://github.com/GoEcosystem/go-web-scraper" class="btn">View on GitHub</a>
  <a href="{{ site.baseurl }}/guide/" class="btn">Explore Documentation</a>
</div>

<div class="example-terminal">
  <div class="terminal-header">
    <span class="terminal-button red"></span>
    <span class="terminal-button yellow"></span>
    <span class="terminal-button green"></span>
    <span class="terminal-title">example terminal</span>
  </div>
  <div class="terminal-content">
    <div class="command">git clone https://github.com/GoEcosystem/go-web-scraper.git</div>
    <div class="command">cd go-web-scraper</div>
    <div class="command">go build ./cmd/scraper</div>
    <div class="command">./scraper -source hackernews -output data/articles.json</div>
    <div class="output">Successfully scraped 30 articles from Hacker News</div>
  </div>
</div>

<h2 class="section-title" id="features">Features</h2>

<div class="card-grid">
  <div class="card">
    <div class="card-body">
      <i class="bi bi-activity card-icon" style="color: var(--primary-color);"></i>
      <h3>Concurrent Scraping</h3>
      <p>Utilize Go's concurrency features for efficient data extraction from multiple sources simultaneously.</p>
      <a href="{{ site.baseurl }}/architecture/" class="card-link">Learn more →</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-database card-icon" style="color: var(--primary-color);"></i>
      <h3>SQLite Integration</h3>
      <p>Persistent storage of scraped data with a robust SQLite database backend.</p>
      <a href="{{ site.baseurl }}/api/" class="card-link">Learn more →</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-window card-icon" style="color: var(--primary-color);"></i>
      <h3>Web Interface</h3>
      <p>Browse and search scraped data through a clean, responsive web UI.</p>
      <a href="{{ site.baseurl }}/guide/" class="card-link">Learn more →</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-file-earmark-arrow-down card-icon" style="color: var(--primary-color);"></i>
      <h3>Multiple Export Formats</h3>
      <p>Export data to JSON or CSV formats for further analysis.</p>
      <a href="{{ site.baseurl }}/api/" class="card-link">Learn more →</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-speedometer2 card-icon" style="color: var(--primary-color);"></i>
      <h3>Rate Limiting</h3>
      <p>Built-in rate limiting to ensure ethical scraping practices.</p>
      <a href="{{ site.baseurl }}/guide/" class="card-link">Learn more →</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-boxes card-icon" style="color: var(--primary-color);"></i>
      <h3>Extensible Design</h3>
      <p>Easily add new scraper modules for additional data sources.</p>
      <a href="{{ site.baseurl }}/api/" class="card-link">Learn more →</a>
    </div>
  </div>
</div>

<h2 class="section-title" id="documentation">Documentation</h2>

<div class="card-grid">
  <div class="card">
    <div class="card-body">
      <i class="bi bi-journal-code card-icon" style="color: var(--primary-color);"></i>
      <h3>API Reference</h3>
      <p>Comprehensive API documentation for interfacing with the Web Scraper.</p>
      <a href="{{ site.baseurl }}/api/" class="card-link">View API Docs</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-diagram-3 card-icon" style="color: var(--primary-color);"></i>
      <h3>Architecture</h3>
      <p>Understand the internal structure and design of the application.</p>
      <a href="{{ site.baseurl }}/architecture/" class="card-link">View Architecture</a>
    </div>
  </div>
  
  <div class="card">
    <div class="card-body">
      <i class="bi bi-book card-icon" style="color: var(--primary-color);"></i>
      <h3>Installation Guide</h3>
      <p>Instructions for setting up and running the web scraper.</p>
      <a href="{{ site.baseurl }}/guide/" class="card-link">View Guide</a>
    </div>
  </div>
</div>

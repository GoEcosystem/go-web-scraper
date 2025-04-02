---
layout: api
title: Go Web Scraper API Reference
description: Complete API documentation for the Go Web Scraper
repo_name: go-web-scraper
breadcrumb_name: API Reference
order: 10
---

{% include breadcrumbs.html %}

<div class="api-documentation">
  <div class="api-sidebar">
    {% include sidebar.html title="API Sections" items=site.data.api_navigation %}
  </div>
  
  <div class="api-content">
    <h1>Go Web Scraper API Reference</h1>
    
    <p>The Go Web Scraper provides a RESTful API for accessing and manipulating scraped data programmatically. This documentation covers all the available endpoints and how to use them.</p>
    
    <h2 id="base-url">Base URL</h2>
    
    <p>When running locally:</p>
    <pre><code>http://localhost:8080/api</code></pre>
    
    <h2 id="authentication">Authentication</h2>
    
    <p>Currently, the API doesn't require authentication as it's designed for local use.</p>
    
    <h2 id="endpoints">Endpoints</h2>
    
    <h3 id="articles">Articles</h3>
    
    <table>
      <thead>
        <tr>
          <th>Method</th>
          <th>Endpoint</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/articles</code></td>
          <td>List all articles with pagination</td>
        </tr>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/articles/:id</code></td>
          <td>Get a specific article by ID</td>
        </tr>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/articles/search</code></td>
          <td>Search articles by keywords</td>
        </tr>
      </tbody>
    </table>
    
    <h4 id="list-articles">List Articles</h4>
    
    <pre><code>GET /api/articles?page=1&limit=20</code></pre>
    
    <p><strong>Query Parameters:</strong></p>
    <ul>
      <li><code>page</code>: Page number (default: 1)</li>
      <li><code>limit</code>: Number of items per page (default: 20)</li>
      <li><code>sort</code>: Sort field (options: 'date', 'title', 'score')</li>
      <li><code>order</code>: Sort order (options: 'asc', 'desc')</li>
    </ul>
    
    <p><strong>Response:</strong></p>
    <pre><code class="language-json">{
  "data": [
    {
      "id": 1,
      "title": "Introducing Go 2.0",
      "url": "https://example.com/go-2-release",
      "score": 120,
      "comments": 45,
      "author": "gopher",
      "date": "2025-03-15T14:22:18Z"
    },
    {
      "id": 2,
      "title": "Go Concurrency Patterns",
      "url": "https://example.com/go-concurrency",
      "score": 98,
      "comments": 32,
      "author": "rob_pike",
      "date": "2025-03-14T09:45:10Z"
    }
  ],
  "meta": {
    "current_page": 1,
    "per_page": 20,
    "total_items": 235,
    "total_pages": 12
  }
}</code></pre>
    
    <h4 id="get-article">Get Article by ID</h4>
    
    <pre><code>GET /api/articles/42</code></pre>
    
    <p><strong>Response:</strong></p>
    <pre><code class="language-json">{
  "id": 42,
  "title": "Understanding Go Interfaces",
  "url": "https://example.com/go-interfaces",
  "score": 156,
  "comments": 67,
  "author": "go_enthusiast",
  "date": "2025-03-10T11:32:45Z",
  "content": "Go interfaces are a powerful feature that enables..."
}</code></pre>
    
    <h4 id="search-articles">Search Articles</h4>
    
    <pre><code>GET /api/articles/search?q=golang+concurrency</code></pre>
    
    <p><strong>Query Parameters:</strong></p>
    <ul>
      <li><code>q</code>: Search query (required)</li>
      <li><code>page</code>: Page number (default: 1)</li>
      <li><code>limit</code>: Number of items per page (default: 20)</li>
    </ul>
    
    <p><strong>Response:</strong> Same format as List Articles</p>
    
    <h3 id="products">Products</h3>
    
    <table>
      <thead>
        <tr>
          <th>Method</th>
          <th>Endpoint</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/products</code></td>
          <td>List all products with pagination</td>
        </tr>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/products/:id</code></td>
          <td>Get a specific product by ID</td>
        </tr>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/products/search</code></td>
          <td>Search products by keywords</td>
        </tr>
      </tbody>
    </table>
    
    <p>Similar patterns apply for other product endpoints.</p>
    
    <h3 id="export">Export</h3>
    
    <table>
      <thead>
        <tr>
          <th>Method</th>
          <th>Endpoint</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/export/articles</code></td>
          <td>Export all articles as JSON or CSV</td>
        </tr>
        <tr>
          <td><span class="endpoint-method get">GET</span></td>
          <td><code>/api/export/products</code></td>
          <td>Export all products as JSON or CSV</td>
        </tr>
      </tbody>
    </table>
    
    <h4 id="export-data">Export Data</h4>
    
    <pre><code>GET /api/export/articles?format=json</code></pre>
    
    <p><strong>Query Parameters:</strong></p>
    <ul>
      <li><code>format</code>: Output format (options: 'json', 'csv', default: 'json')</li>
    </ul>
    
    <h2 id="error-handling">Error Handling</h2>
    
    <p>All API errors return an appropriate HTTP status code and a JSON error response:</p>
    
    <pre><code class="language-json">{
  "error": {
    "code": "not_found",
    "message": "Article with ID 999 not found"
  }
}</code></pre>
    
    <h3 id="common-error-codes">Common Error Codes</h3>
    
    <table>
      <thead>
        <tr>
          <th>HTTP Status</th>
          <th>Error Code</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>400</td>
          <td>invalid_request</td>
          <td>The request is malformed or missing required parameters</td>
        </tr>
        <tr>
          <td>404</td>
          <td>not_found</td>
          <td>The requested resource was not found</td>
        </tr>
        <tr>
          <td>422</td>
          <td>validation_error</td>
          <td>Validation errors in the request parameters</td>
        </tr>
        <tr>
          <td>500</td>
          <td>internal_error</td>
          <td>An unexpected error occurred on the server</td>
        </tr>
      </tbody>
    </table>
  </div>
</div>

<style>
  .api-documentation {
    display: flex;
    gap: 2rem;
  }
  
  .api-sidebar {
    width: 250px;
    flex-shrink: 0;
  }
  
  .api-content {
    flex: 1;
  }
  
  @media (max-width: 768px) {
    .api-documentation {
      flex-direction: column;
    }
    
    .api-sidebar {
      width: 100%;
      margin-bottom: 2rem;
    }
  }
  
  /* Endpoint method styling */
  .endpoint-method {
    display: inline-block;
    padding: 0.25rem 0.5rem;
    border-radius: 4px;
    font-weight: bold;
    margin-right: 0.5rem;
    font-size: 0.9rem;
  }
  
  .endpoint-method.get {
    background-color: #D1ECF1;
    color: #0C5460;
  }
  
  .endpoint-method.post {
    background-color: #D4EDDA;
    color: #155724;
  }
  
  .endpoint-method.put {
    background-color: #FFF3CD;
    color: #856404;
  }
  
  .endpoint-method.delete {
    background-color: #F8D7DA;
    color: #721C24;
  }
</style>

---
layout: api
title: API Documentation
---

# API Documentation

The Go Web Scraper provides a RESTful API for accessing and manipulating scraped data programmatically. This documentation covers all the available endpoints and how to use them.

## Base URL

When running locally:
```
http://localhost:8080/api
```

## Authentication

Currently, the API doesn't require authentication as it's designed for local use.

## Endpoints

### Articles

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/articles` | List all articles with pagination |
| GET | `/api/articles/:id` | Get a specific article by ID |
| GET | `/api/articles/search` | Search articles by keywords |

#### List Articles

```
GET /api/articles?page=1&limit=20
```

Query Parameters:
- `page`: Page number (default: 1)
- `limit`: Number of items per page (default: 20)
- `sort`: Sort field (options: 'date', 'title', 'score')
- `order`: Sort order (options: 'asc', 'desc')

Response:
```json
{
  "data": [
    {
      "id": 1,
      "title": "Introducing Go 2.0",
      "url": "https://example.com/go-2-release",
      "score": 342,
      "comments": 127,
      "date": "2025-03-30T15:42:31Z",
      "author": "gopher"
    },
    ...
  ],
  "pagination": {
    "total": 256,
    "page": 1,
    "limit": 20,
    "totalPages": 13
  }
}
```

#### Get Article by ID

```
GET /api/articles/42
```

Response:
```json
{
  "id": 42,
  "title": "Advanced Concurrency Patterns in Go",
  "url": "https://example.com/concurrency-patterns",
  "score": 512,
  "comments": 89,
  "date": "2025-03-25T09:17:22Z",
  "author": "concurrency_fan",
  "content": "Full article content here...",
  "tags": ["go", "concurrency", "programming"]
}
```

#### Search Articles

```
GET /api/articles/search?q=golang+concurrency
```

Query Parameters:
- `q`: Search query
- `page`: Page number (default: 1)
- `limit`: Number of items per page (default: 20)

Response: Same format as List Articles.

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products` | List all products with pagination |
| GET | `/api/products/:id` | Get a specific product by ID |
| GET | `/api/products/search` | Search products by keywords |
| GET | `/api/products/category/:category` | List products by category |

#### List Products

```
GET /api/products?page=1&limit=20
```

Query Parameters:
- `page`: Page number (default: 1)
- `limit`: Number of items per page (default: 20)
- `sort`: Sort field (options: 'price', 'title', 'rating')
- `order`: Sort order (options: 'asc', 'desc')

Response:
```json
{
  "data": [
    {
      "id": 1,
      "title": "Programming in Go",
      "author": "Mark Author",
      "price": 39.99,
      "imageUrl": "https://example.com/book-cover.jpg",
      "rating": 4.7,
      "category": "programming"
    },
    ...
  ],
  "pagination": {
    "total": 178,
    "page": 1,
    "limit": 20,
    "totalPages": 9
  }
}
```

Similar patterns apply for other product endpoints.

### Export

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/export/articles` | Export all articles as JSON or CSV |
| GET | `/api/export/products` | Export all products as JSON or CSV |

#### Export Articles

```
GET /api/export/articles?format=json
```

Query Parameters:
- `format`: Export format (options: 'json', 'csv', default: 'json')

Response: File download

## Status Codes

The API uses standard HTTP status codes:

- `200 OK`: Request successful
- `400 Bad Request`: Invalid request parameters
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Rate Limiting

To prevent abuse, the API implements rate limiting of 100 requests per minute per IP address. If you exceed this limit, you'll receive a `429 Too Many Requests` response.

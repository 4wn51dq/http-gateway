# Simple Product Accessing API

A simple **Go REST API** built using the standard `net/http` package.  

The API is documented using **Swagger / OpenAPI**, allowing interactive exploration and client generation.

---

## Purpose:

help understand:
- how REST APIs work internally in Go and structure the backend with `net/http`
- designing and exposing API endpoints 
- How Swagger/OpenAPI documents an API
- How API contracts differ from implementation

---

- **Language:** Go
- **HTTP Server:** `net/http`
- **API Style:** REST
- **API Documentation:** Swagger / OpenAPI
- **Data Storage:** In-memory (for simplicity)

---

## Features

- Create, read, update, and delete products
- RESTful endpoint design
- Swagger UI for API exploration
- Clear separation between routing, handlers, and models
- Wrapped the server's servemux handler with gohandler.CORS for enabling CORS

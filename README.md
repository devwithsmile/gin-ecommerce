# ShopSphere Backend (Go + Gin + Elasticsearch)

Backend for **ShopSphere**, a modern e‑commerce platform designed around fast search and scalable architecture. The backend is written in Go using the Gin framework and is intended to use Elasticsearch as the primary data store (currently starting with in‑memory storage for the customer/auth slice).[file:77][web:48]

---

## Architecture Overview

**High‑level design**

- **API layer**: Gin HTTP server exposes REST endpoints for customers, sellers, and admins (starting with customer auth).
- **Auth layer**: JWT‑based authentication and authorization using `github.com/golang-jwt/jwt/v5` with access and refresh tokens.
- **Domain layer**: Separate packages per domain (e.g. `customer`) containing models, services, and repositories.
- **Persistence layer**: 
  - Current: in‑memory repositories for fast iteration.
  - Planned: Elasticsearch for products, orders, and analytics as described in the product requirements document.

This structure is inspired by common Go project layout and clean architecture style, keeping HTTP transport, domain logic, and infrastructure concerns separated.

---

## Tech Stack

- **Language**: Go
- **Web framework**: [Gin](https://github.com/gin-gonic/gin)
- **Planned datastore**: Elasticsearch for product catalog, search, and analytics.



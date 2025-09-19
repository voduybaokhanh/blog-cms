# Blog CMS API (Go: Gin + GORM)

A minimal CMS API built with Go, Gin, and GORM.

- CRUD for Users, Posts, Categories, Tags
- JWT Authentication
- Role-based Access Control
- MySQL Database

## Table of Contents

1. Features
2. Tech Stack
3. Project Structure
4. Configuration (.env)
5. Installation & Run
6. API Endpoints
7. JWT Usage
8. Postman Collection
9. Notes
10. Author

---

## 1) Features

- Users CRUD (admin protected)
- Posts CRUD (auth required; author/admin protections)
- Categories & Tags CRUD (admin protected)
- JWT-based authentication
- Role-based access control (RBAC)
- MySQL database integration

## 2) Tech Stack

- Go >= 1.20 — https://go.dev/
- Gin (Web Framework) — https://github.com/gin-gonic/gin
- GORM (ORM) — https://gorm.io/
- MySQL (Database) — https://www.mysql.com/
- golang-jwt (JWT) — https://github.com/golang-jwt/jwt
- Postman (API testing) — https://www.postman.com/

## 3) Project Structure

```text
blog-cms/
│── cmd/           # Entry point
│   └── main.go
│── controllers/   # Controllers (Auth, User, Post, Category, Tag)
│── middleware/    # Auth, roles middleware
│── models/        # DB models
│── pkg/           # Utilities (hash, jwt)
│── routes/        # API routes
│── config/        # DB config, load env
│── blog-cms.postman_collection.json
│── go.mod
│── go.sum
│── README.md
```

## 4) Configuration (.env)

Create a `.env` file in the project root:

```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=blogcms
JWT_SECRET=supersecretkey
PORT=3000
```

## 5) Installation & Run

1. Clone project
   ```bash
   git clone https://github.com/voduybaokhanh/blog-cms.git
   cd blog-cms
   ```
2. Install dependencies
   ```bash
   go mod tidy
   ```
3. Create MySQL database
   ```sql
   CREATE DATABASE blogcms;
   ```
4. Run the project
   ```bash
   go run cmd/main.go
   ```

Server will start at: http://localhost:3000

## 6) API Endpoints

### Auth

- POST `/api/v1/auth/register` — Register
- POST `/api/v1/auth/login` — Login (returns JWT token)
- GET `/api/v1/me` — Current user (requires Bearer token)

### Users (admin only group)

- GET `/api/v1/users`
- GET `/api/v1/users/:id`
- PUT `/api/v1/users/:id`
- DELETE `/api/v1/users/:id`

### Posts (auth required group)

- GET `/api/v1/posts`
- GET `/api/v1/posts/:id`
- POST `/api/v1/posts`
- PUT `/api/v1/posts/:id` (author or admin)
- DELETE `/api/v1/posts/:id` (author or admin)

### Categories (admin only group)

- GET `/api/v1/categories`
- POST `/api/v1/categories`
- DELETE `/api/v1/categories/:id`

### Tags (admin only group)

- GET `/api/v1/tags`
- POST `/api/v1/tags`
- DELETE `/api/v1/tags/:id`

## 7) JWT Usage

Obtain a token via `POST /api/v1/auth/login`, then include it in subsequent requests:

```http
Authorization: Bearer <token>
```

Example:

```http
GET /api/v1/me
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
```

## 8) Postman Collection

Import the collection file in the project root:

```
blog-cms.postman_collection.json
```

Notes:
- Collection defines `base_url` and stores `token` in collection variables.
- After login, token is auto-saved; subsequent requests use Bearer auth automatically.

## 9) Notes

- Default port is `3000` (configurable in `.env`)
- If MySQL connection fails, verify `.env` and your MySQL user/password
- JWT token default expiration is 24h

---

✅ Tiến độ (theo tuần)

Week 1: Setup dự án, Auth (Register/Login/Me) (Done)

Week 2: CRUD Users + JWT Middleware + Role-based Access

Week 3: CRUD Blogs + Pagination + Search

Week 4: Testing với Postman + Deployment (Docker Compose)

## 10) Author

- voduybaokhanh

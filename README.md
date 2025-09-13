# Mini CMS API (Go: Gin + GORM)

A minimal CMS API built with Go, Gin, and GORM.

- CRUD for Users and Blogs
- JWT Authentication
- Role-based Access Control
- Pagination & Search
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
11. Vietnamese (Original Content)

---

## 1) Features

- User and Blog CRUD
- JWT-based authentication
- Role-based access control (RBAC)
- Pagination and search for blogs
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
│── controllers/   # Controllers (Auth, User, Blog)
│── middleware/    # JWT middleware
│── models/        # DB models
│── pkg/           # Shared packages (hash password, utils)
│── routes/        # API routes
│── config/        # DB config, load env
│── go.mod
│── go.sum
│── .env           # Environment configuration
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
   go run ./cmd
   ```

Server will start at: http://localhost:3000

## 6) API Endpoints

### Auth

- POST `/api/v1/auth/register` — Register
- POST `/api/v1/auth/login` — Login (returns JWT token)
- GET `/api/v1/me` — Get current user info (requires Bearer token)

### Users

- GET `/api/v1/users` — List users (Admin)
- GET `/api/v1/users/:id` — User details
- PUT `/api/v1/users/:id` — Update user
- DELETE `/api/v1/users/:id` — Delete user

### Blogs

- GET `/api/v1/blogs` — List blogs (supports search, page, limit)
- GET `/api/v1/blogs/:id` — Blog details
- POST `/api/v1/blogs` — Create a blog (requires login)
- PUT `/api/v1/blogs/:id` — Update a blog
- DELETE `/api/v1/blogs/:id` — Delete a blog

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

Postman collection is available at:

```bash
/postman/blog-cms.postman_collection.json
```

How to import:

1) Open Postman
2) File → Import
3) Choose `blog-cms.postman_collection.json`

## 9) Notes

- Default port is `3000` (configurable in `.env`)
- If MySQL connection fails, verify `.env` and your MySQL user/password
- JWT token default expiration is 24h
- AUTO_INCREMENT is not reset when deleting users (keeps ID integrity)

## 10) Author

- voduybaokhanh

---

✅ Tiến độ (theo tuần)

Week 1: Setup dự án, Auth (Register/Login/Me) (Done

Week 2: CRUD Users + JWT Middleware + Role-based Access

Week 3: CRUD Blogs + Pagination + Search

Week 4: Testing với Postman + Deployment (Docker Compose)

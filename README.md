# 🚀 Blog CMS API

<div align="center">

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-Database-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)

*A modern, minimal CMS API built with Go, Gin, and GORM*

[Features](#-features) • [Tech Stack](#-tech-stack) • [Installation](#-installation) • [API Documentation](#-api-documentation) • [Usage](#-usage)

</div>

---

## ✨ Features

### 🔐 Authentication & Authorization
- **JWT-based authentication** with secure token management
- **Role-based access control (RBAC)** for fine-grained permissions
- **Protected routes** with middleware validation

### 📝 Content Management
- **Users CRUD** (admin protected)
- **Posts CRUD** with author/admin protection
- **Categories & Tags CRUD** (admin protected)
- **Advanced search** across title and content
- **Smart filtering** by categories and tags

### 🔍 Advanced Querying
- **Pagination** for efficient data loading
- **Search functionality** with full-text capabilities
- **Multi-tag filtering** support
- **Category-based filtering**

### 🗄️ Database
- **MySQL integration** with GORM ORM
- **Optimized queries** for better performance
- **Relationship management** between entities

---

## 🛠️ Tech Stack

| Technology | Version | Purpose |
|------------|---------|---------|
| [**Go**](https://go.dev/) | ≥ 1.20 | Programming language |
| [**Gin**](https://github.com/gin-gonic/gin) | Latest | Web framework |
| [**GORM**](https://gorm.io/) | Latest | ORM library |
| [**MySQL**](https://www.mysql.com/) | 8.0+ | Database |
| [**golang-jwt**](https://github.com/golang-jwt/jwt) | Latest | JWT authentication |
| [**Postman**](https://www.postman.com/) | Latest | API testing |

---

## 📁 Project Structure

```
blog-cms/
├── 📁 cmd/                    # Application entry point
│   └── main.go
├── 📁 controllers/            # HTTP handlers
│   ├── auth.go
│   ├── user.go
│   └── post.go
├── 📁 middleware/             # Custom middleware
│   ├── auth.go
│   └── roles.go
├── 📁 models/                 # Database models
│   ├── user.go
│   ├── post.go
│   ├── category.go
│   └── tag.go
├── 📁 pkg/                    # Utility packages
│   ├── hash/
│   └── jwt/
├── 📁 routes/                 # API route definitions
├── 📁 config/                 # Configuration management
├── 📄 blog-cms.postman_collection.json
├── 📄 go.mod
├── 📄 go.sum
└── 📄 README.md
```

---

## ⚙️ Configuration

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=blogcms

# JWT Configuration
JWT_SECRET=supersecretkey

# Server Configuration
PORT=3000
```

---

## 🚀 Installation & Setup

### Prerequisites
- Go 1.20 or higher
- MySQL 8.0 or higher
- Git

### Step 1: Clone the Repository
```bash
git clone https://github.com/voduybaokhanh/blog-cms.git
cd blog-cms
```

### Step 2: Install Dependencies
```bash
go mod tidy
```

### Step 3: Database Setup
```sql
CREATE DATABASE blogcms;
```

### Step 4: Configure Environment
```bash
# Copy the example environment file
cp .env.example .env

# Edit the .env file with your database credentials
```

### Step 5: Run the Application
```bash
go run cmd/main.go
```

🎉 **Server will be running at:** `http://localhost:3000`

---

## 📚 API Documentation

### 🔐 Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/api/v1/auth/register` | Register new user | ❌ |
| `POST` | `/api/v1/auth/login` | Login user | ❌ |
| `GET` | `/api/v1/me` | Get current user | ✅ |

### 👥 User Management (Admin Only)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/api/v1/users` | List all users | ✅ Admin |
| `GET` | `/api/v1/users/:id` | Get user by ID | ✅ Admin |
| `PUT` | `/api/v1/users/:id` | Update user | ✅ Admin |
| `DELETE` | `/api/v1/users/:id` | Delete user | ✅ Admin |

### 📝 Post Management

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/api/v1/posts` | List posts (with pagination & filters) | ✅ |
| `GET` | `/api/v1/posts/:id` | Get single post | ✅ |
| `POST` | `/api/v1/posts` | Create new post | ✅ |
| `PUT` | `/api/v1/posts/:id` | Update post | ✅ Author/Admin |
| `DELETE` | `/api/v1/posts/:id` | Delete post | ✅ Author/Admin |

### 🏷️ Category Management (Admin Only)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/api/v1/categories` | List all categories | ✅ Admin |
| `POST` | `/api/v1/categories` | Create category | ✅ Admin |
| `DELETE` | `/api/v1/categories/:id` | Delete category | ✅ Admin |

### 🏷️ Tag Management (Admin Only)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/api/v1/tags` | List all tags | ✅ Admin |
| `POST` | `/api/v1/tags` | Create tag | ✅ Admin |
| `DELETE` | `/api/v1/tags/:id` | Delete tag | ✅ Admin |

### 🔍 Query Parameters for Posts

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Page number for pagination |
| `limit` | integer | 10 | Number of posts per page |
| `search` | string | - | Search in title and content |
| `category` | integer | - | Filter by category ID |
| `tag` | string | - | Filter by tag IDs (comma-separated) |

### 📋 Example Requests

```bash
# Pagination
GET /api/v1/posts?page=1&limit=5

# Search
GET /api/v1/posts?search=golang

# Filter by category
GET /api/v1/posts?category=2

# Filter by multiple tags
GET /api/v1/posts?tag=1,2

# Combined filters
GET /api/v1/posts?search=api&category=1&tag=3,4&page=2&limit=10
```

---

## 🔑 JWT Usage

### Authorization Header
```http
Authorization: Bearer <your-jwt-token>
```

### Example Request
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6..." \
     http://localhost:3000/api/v1/me
```

---

## 🧪 Testing with Postman

1. **Import Collection**: Import `blog-cms.postman_collection.json`
2. **Set Variables**: Configure `base_url` and `token` variables
3. **Auto-token**: Login endpoint automatically saves token for subsequent requests

---

## 📝 Important Notes

- **Default Port**: 3000 (configurable in `.env`)
- **JWT Expiry**: 24 hours (configurable)
- **Database**: Ensure MySQL connection is properly configured
- **Environment**: Check `.env` file if database connection fails

---

## 📊 Development Progress

- ✅ **Week 1**: Project setup, Authentication (Register/Login/Me)
- ✅ **Week 2**: User CRUD, JWT Middleware, RBAC implementation
- ✅ **Week 3**: Post CRUD, Pagination, Search, Post-Tags relationships
- 🔜 **Week 4**: Postman testing, Docker Compose deployment

---

## 👨‍💻 Author

**voduybaokhanh**

---

<div align="center">

**⭐ Star this repository if you found it helpful!**

Made with ❤️ using Go

</div>
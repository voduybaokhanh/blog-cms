# Sử dụng Go image chính thức
FROM golang:1.22 AS builder

# Set thư mục làm việc
WORKDIR /app

# Copy go.mod và go.sum trước (để cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary
RUN go build -o blog-cms ./cmd/main.go

# Stage 2: minimal runtime
FROM debian:bookworm-slim

# Cài các thư viện cần thiết
RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy binary từ stage build
COPY --from=builder /app/blog-cms .

# Copy file .env (nếu có)
COPY .env .env

# Expose port cho Gin
EXPOSE 8080

CMD ["./blog-cms"]

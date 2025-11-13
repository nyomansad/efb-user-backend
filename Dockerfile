# Gunakan base image Golang
FROM golang:1.24 AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependencies
RUN go mod download

# Build binary aplikasi
RUN go build -o main .

# --- Tahap kedua: membuat image lebih kecil ---
FROM debian:bookworm-slim

WORKDIR /root/

# Copy binary dari tahap builder
COPY --from=builder /app/main .

# Port yang digunakan oleh Gin (biasanya 8080)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]

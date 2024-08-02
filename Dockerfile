# Use the official Golang image to build the binary
FROM golang:1.22 as builder

# Set the architecture for the build
ENV GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o api ./cmd/api/main.go
RUN go build -o consumer ./cmd/consumer/main.go
RUN go build -o cron ./cmd/cron/main.go

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/api .
COPY --from=builder /app/consumer .
COPY --from=builder /app/cron .

# Copy the config.yaml file
COPY config.yaml /root/

# Expose ports
EXPOSE 8080

# Default command (can be overridden by docker-compose)
# CMD ["/root/api"]

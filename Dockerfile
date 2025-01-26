# Step 1: Use official Golang image as a base
FROM golang:1.20 AS builder

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the source code
COPY . .

# Step 6: Build the application
RUN go build -o main ./cmd/userService/main.go

# Step 7: Use a minimal base image
FROM debian:bullseye-slim

# Step 8: Set working directory
WORKDIR /app

# Step 9: Copy the binary from the builder
COPY --from=builder /app/main .

# Step 10: Expose the application port
EXPOSE 5000

# Step 11: Command to run the application
CMD ["./main"]

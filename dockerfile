# Stage 1: Build the Go application
FROM golang:1.22.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Stage 2: Run the Go application
FROM golang:1.22.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the previous stage
COPY --from=builder /app .


# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the Go application
CMD ["go", "run", "main.go"]

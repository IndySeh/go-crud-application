# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application with static linking for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crud-app ./cmd/api/main.go

RUN mkdir -p /go/src/logs

# Stage 2: Create a minimal final image using 'scratch'
FROM scratch AS build-release-stage

# Set the working directory inside the final container
WORKDIR /go/src/project

# Copy the statically compiled binary from the 'builder' stage
COPY --from=builder /app/crud-app .
# Expose port 8080 for the Go application
EXPOSE 8080

# Command to run the Go application
CMD ["./crud-app"]
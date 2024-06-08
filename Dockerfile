# Use the official Golang image as the base image
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Set up Go module
ENV GO111MODULE=on

# Install dependencies (if needed)
# RUN go mod download

# Build the Go application
RUN go build -o tanahore ./cmd/api/main.go

# Expose the port that your API will run on
EXPOSE 8080

# Command to run the executable
CMD ["./tanahore"]

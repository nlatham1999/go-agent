# Start from the official Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go server
RUN go build -o main .

# Expose the port that your server will run on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]

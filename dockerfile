# Use an official Go runtime as the base image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /app

# Copy the Go module and source code into the container
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN go build -o main .

EXPOSE 8080
# Expose the port on which the application will run

# Command to run the application
CMD ["./main"]
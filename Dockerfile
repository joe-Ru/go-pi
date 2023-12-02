# Use an official Go runtime as a parent image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o go_pi .

# Expose the port the application runs on
EXPOSE 8080

# Run the application
CMD ["./go_pi"]
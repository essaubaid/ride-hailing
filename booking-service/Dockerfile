# Use a lightweight Go base image.
FROM golang:1.23.4

# Set the working directory in the container.
WORKDIR /app

# Copy the Go modules and dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go application.
RUN go build -o booking-service ./booking-service

# Expose the port your service listens on.
EXPOSE 80

# Run the executable.
CMD ["./booking-service/booking-service"]

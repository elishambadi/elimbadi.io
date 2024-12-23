# Step 1: Use the official Golang image to build your app
FROM golang:1.20-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go modules and download dependencies
COPY go.mod ./
RUN go mod tidy

# Step 4: Copy the application code into the container
COPY . .

# Step 5: Build the Go app from the cmd/server directory
WORKDIR /app/cmd/server
RUN go build -o /usr/local/bin/myapp .

# Step 6: Use a smaller image for the final image to run the app
FROM alpine:latest

# Step 7: Copy the binary from the builder image
COPY --from=builder /usr/local/bin/myapp /usr/local/bin/myapp

# Step 8: Expose the port your app runs on
EXPOSE 8000

# Step 9: Run the application
CMD ["myapp"]

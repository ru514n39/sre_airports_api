# Build stage
FROM golang:1.23.1 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and source code
COPY go.mod ./
RUN go mod tidy

# Install required dependencies
RUN go get github.com/aws/aws-sdk-go/aws \
           github.com/aws/aws-sdk-go/aws/session \
           github.com/aws/aws-sdk-go/service/s3

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Final stage - use a minimal image
FROM ubuntu

# Set the working directory
WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./main"]

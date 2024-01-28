# Start from the latest golang base image
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
# Make sure to copy the cmd directory and any other directories or files you need
COPY services/user-service ./services/user-service

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./services/user-service/cmd

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/user-service .

# Command to run the executable
CMD ["./user-service"]

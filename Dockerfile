FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .




RUN go build -o main ./cmd/phonedb

# Final stage
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/number.txt .


# Command to run the executable
CMD ["./main"]
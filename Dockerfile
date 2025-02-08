FROM golang:1.23.5-alpine as builder

WORKDIR /app/kitchen-sink-server

# Copy the Go module files
COPY go.mod .
COPY go.sum .
COPY db/migrations migrations

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN GOOS=linux go build -o go-server .

FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=builder /app/kitchen-sink-server /app/kitchen-sink-server

WORKDIR /app/kitchen-sink-server
EXPOSE 8080
ENTRYPOINT ["/app/kitchen-sink-server/go-server"]

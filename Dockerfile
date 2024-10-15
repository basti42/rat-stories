#  Development
FROM golang:1.23 as dev

WORKDIR /server

# Build
FROM golang:1.23 as build
WORKDIR /server

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux go build -v -o / /server/...

# Production
FROM debian:bookworm-slim as prod
WORKDIR /

COPY --from=build /stories-service /stories-service

# Install tls certificates
RUN apt-get update && apt-get install -y ca-certificates

CMD ["/stories-service"]
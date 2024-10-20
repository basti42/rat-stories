# Build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# cgo enabled must be true, do to sqlite requirement
RUN CGO_ENABLED=1 GOOS=linux go build -a --installsuffix cgo -o main .

# Production
FROM alpine:3.20 AS prod

# install necessary certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]

# Build stage
FROM golang:1.24.2-alpine AS builder
WORKDIR /src

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go-ci-app ./cmd

# Runtime stage
FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /go-ci-app /usr/local/bin/go-ci-app
USER app
EXPOSE 8080
ENV APP_VERSION=dev
ENTRYPOINT ["/usr/local/bin/go-ci-app"]

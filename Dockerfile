FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o teste ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/teste .

ENTRYPOINT ["/app/teste"]

FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /beat-exporter .

FROM alpine:latest
RUN adduser -D -u 10001 exporter
COPY --from=builder /beat-exporter /bin/beat-exporter
USER exporter
EXPOSE 9479
ENTRYPOINT ["/bin/beat-exporter"]

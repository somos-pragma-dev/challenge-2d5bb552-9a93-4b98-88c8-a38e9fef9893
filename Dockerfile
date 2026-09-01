FROM golang:1.22 AS builder
WORKDIR /app
COPY..
RUN go build -o app

FROM alpine:latest
COPY --from=builder /app/app /app/app
EXPOSE 50051
CMD ["./app"]
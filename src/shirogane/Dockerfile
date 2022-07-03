FROM golang:1.18-alpine as builder
WORKDIR /app/
COPY . .
RUN go build -o /app/shirogane cmd/main.go

FROM alpine:3.6
WORKDIR /app/
COPY --from=builder ["/app/shirogane", "/app/.env", "./"]
EXPOSE 8000
CMD /app/shirogane

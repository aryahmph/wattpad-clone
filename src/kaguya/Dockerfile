FROM golang:1.18-alpine as builder
WORKDIR /app/
COPY . .
RUN go build -o /app/kaguya cmd/main.go

FROM alpine:3.6
WORKDIR /app/
COPY --from=builder ["/app/kaguya", "/app/.env", "/app/config/","./"]
EXPOSE 8010
CMD /app/kaguya

FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o ./cmd/api

EXPOSE 8080

CMD ["./go-tanahore.exe"]

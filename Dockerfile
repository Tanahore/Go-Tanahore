FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o tanahore.exe ./cmd/api/main.go

EXPOSE 8080

CMD ["./tanahore.exe"]

FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o go-tanahore.exe ./cmd/api

EXPOSE 8080

CMD ["./go-tanahore"]

FROM golang:1.20-bookworm

WORKDIR /app

COPY . .

RUN go build -o go-tanahore.exe ./cmd/api

EXPOSE 8080

CMD ["./go-tanahore"]

FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o go-tanahore .

EXPOSE 8080

CMD ["./go-tanahore"]

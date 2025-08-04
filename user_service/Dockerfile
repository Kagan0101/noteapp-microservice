# docker-compose.yml ile aynı klasördeyseniz:

FROM golang:1.24.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . # Tüm klasörü kopyala

WORKDIR /app/user_service

RUN go build -o main .

CMD ["./main"]

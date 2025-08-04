

FROM golang:1.24.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . # Tüm klasörü kopyala

WORKDIR /app/note_service

RUN go build -o main .

CMD ["./main"]
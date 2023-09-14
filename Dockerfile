
FROM golang:1.20 as dev

RUN go install github.com/cosmtrek/air@latest


WORKDIR /app

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM golang:1.18

WORKDIR /app

COPY ./services/xwing .

RUN go mod download

CMD ["go", "run", "main.go"]
FROM golang:1.18 AS builder

WORKDIR /app

COPY ./services/deathstar/go.mod ./services/deathstar/go.sum ./
RUN go mod download

COPY ./services/deathstar/ ./

RUN go build -o deathstar .

FROM gcr.io/distroless/base

COPY --from=builder /app/deathstar /deathstar

CMD ["/deathstar"]
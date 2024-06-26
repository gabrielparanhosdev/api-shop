FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM frolvlad/alpine-glibc

COPY --from=builder /app/main /

EXPOSE 4000

CMD ["/main"]

FROM golang:latest

WORKDIR /app

COPY . .

RUN go build servidor.go
RUN go build cliente.go

CMD ["./servidor"]
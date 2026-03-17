FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app ./cmd/main.go

EXPOSE 8081

CMD ["./app"]
FROM golang:1.23.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY key.json .

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]

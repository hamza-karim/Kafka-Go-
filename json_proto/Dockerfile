FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y kafkacat

RUN go build -o main .

CMD ["./main"]


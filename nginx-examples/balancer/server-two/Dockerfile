FROM golang:latest


WORKDIR /app

COPY . .

RUN go mod init balancer

RUN go build -o app

CMD ["./app"]
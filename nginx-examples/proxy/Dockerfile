FROM golang:latest

WORKDIR /app

COPY . .

EXPOSE 3000
RUN go mod init test


RUN go build -o app

CMD ["./app"]
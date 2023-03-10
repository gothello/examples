FROM golang

WORKDIR /app

COPY . .

# RUN go mod donwload

RUN go build -o appgo .

CMD ["./appgo"]
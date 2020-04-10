FROM golang:1.14.2-alpine3.11

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download

RUN go build -o main .

CMD ["/app/main"]

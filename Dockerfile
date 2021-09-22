FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 5000

RUN go mod tidy

RUN swag init

RUN go build -o main

CMD ["/app/main"]

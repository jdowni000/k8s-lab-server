FROM golang:latest

ADD . /app

WORKDIR /app

RUN go build -v

ENTRYPOINT [ "/app/app" ]




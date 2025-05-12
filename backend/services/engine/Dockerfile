FROM golang:latest

RUN apt-get update

RUN mkdir /app

ADD . /app
WORKDIR /app

# Build app
RUN go build -o ./bin/engine ./cmd
CMD ["/app/bin/engine"]

FROM golang:buster

COPY . /app
WORKDIR /app

RUN export GOPATH=$HOME/go
RUN export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
RUN go mod download
RUN go build -o bin/appolo-scraper *.go

ENTRYPOINT [ "./bin/appolo-scraper" ]

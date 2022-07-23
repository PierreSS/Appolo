FROM golang:buster

COPY . /app
WORKDIR /app

RUN export GOPATH=$HOME/go
RUN export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
RUN go install github.com/swaggo/swag/cmd/swag
RUN go mod download
RUN make

ENTRYPOINT [ "./bin/appolo-api" ]

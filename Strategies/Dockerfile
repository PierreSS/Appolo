FROM golang:buster

COPY . /app
WORKDIR /app

#? Install Golang and deps
RUN export GOPATH=$HOME/go
RUN export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
RUN go mod download

#? Install Python and deps
RUN apt update
RUN apt upgrade -y
RUN apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev libsqlite3-dev wget libbz2-dev -y
RUN wget https://www.python.org/ftp/python/3.10.0/Python-3.10.0.tgz
RUN tar -xf Python-3.10.*.tgz
RUN Python-3.10.0/configure --enable-optimizations
RUN make altinstall

# For Ubuntu
# RUN apt install software-properties-common -y
# RUN add-apt-repository ppa:deadsnakes/ppa -y
# RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys BA6932366A755776
# RUN apt-get update
# RUN apt install python3.10

RUN apt-get -y install python3-pip
RUN python3.10 -m pip install --upgrade pip
RUN python3.10 -m pip install -r requirements.txt

RUN go build -o bin/appolo-strat main.go

ENTRYPOINT [ "./bin/appolo-strat" ]

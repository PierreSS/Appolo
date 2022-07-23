all: go

exec = appolo-api

go: main.go
	go build -o bin/$(exec) main.go
	go generate

clean:
	rm -f bin/$(exec) *~ *#
	rm -rf src/gopkg.in
	go mod tidy

re: clean all

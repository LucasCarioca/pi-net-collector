.PHONY := all

start:
	go run .

install:
	go get -d ./...
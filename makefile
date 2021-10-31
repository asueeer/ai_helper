.PHONY: build run

build:
	go build nearby

run:
	go build nearby && ./nearby
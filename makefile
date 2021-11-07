.PHONY: build run

build:
	go build ai_helper

run:
	go build ai_helper && ./ai_helper
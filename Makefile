.DEFAULT_GOAL := goapp

.PHONY: chmod
chmod:
	chmod +x goapp/scripts/run_tests.sh

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*

.PHONY: client
client:
	mkdir -p bin
	go build -o bin/client ./cmd/client
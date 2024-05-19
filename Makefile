.DEFAULT_GOAL := goapp

.PHONY: chmod
chmod:
	echo "Setting executable permissions..."
	chmod +x scripts/run_tests.sh
	chmod +x scripts/run_client.sh

.PHONY: all
all: clean chmod goapp client

.PHONY: goapp
goapp: chmod client
	mkdir -p bin
	go build -o bin ./...

.PHONY: client
client:
	mkdir -p bin
	go build -o bin/client ./cmd/client

.PHONY: clean
clean:
	go clean
	rm -f bin/*
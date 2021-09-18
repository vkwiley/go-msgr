all: build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o bin/msgr cmd/main.go

.PHONY: clean
clean:
	rm -f bin/msgr

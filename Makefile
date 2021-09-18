all: build

.PHONY: test
test:
	go test ./... -v

.PHONY: build
build:
	go build -o bin/msgr cmd/main.go

.PHONY: clean
clean:
	rm -f bin/msgr

.PHONY: fmt
fmt:
	go fmt ./...

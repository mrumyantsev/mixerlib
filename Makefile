.PHONY: build
build:
	@go build -o ./build/passgen ./cmd/passgen/main.go

.PHONY: run
run:
	@go run ./cmd/passgen/main.go

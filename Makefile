.PHONY: build
build:
	@go build -o ./build/unguess ./cmd/unguess/main.go

.PHONY: run
run:
	@go run ./cmd/unguess/main.go

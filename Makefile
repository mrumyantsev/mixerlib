.SILENT:
.DEFAULT_GOAL := fast-run

.PHONY: build
build:
	go build -o ./build/unguess ./cmd/unguess/main.go

.PHONY: run
run:
	./build/unguess

.PHONY: fast-run
fast-run:
	go run ./cmd/unguess/main.go

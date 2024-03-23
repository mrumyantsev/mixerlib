.SILENT:
.DEFAULT_GOAL := fast-run

.PHONY: build
build:
	go build -o ./build/unguess ./cmd/unguess

.PHONY: run
run:
	./build/unguess

.PHONY: fast-run
fast-run:
	go run ./cmd/unguess

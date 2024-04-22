.SILENT:
.DEFAULT_GOAL := fast-run

.PHONY: build
build:
	go build -o ./build/mixer ./cmd/mixer

.PHONY: run
run:
	./build/mixer

.PHONY: fast-run
fast-run:
	go run ./cmd/mixer

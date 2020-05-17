test:
	go test .

build:
	go build .

bench:
	go test -bench=. .

run:
	go run .

.PHONY: test build bench

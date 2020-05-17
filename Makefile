test:
	go test .

build:
	go build .

clean:
	rm -f yagolisp

bench:
	go test -bench=. .

run:
	go run .

.PHONY: test build clean bench run

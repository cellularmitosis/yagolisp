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

types: build
	./yagolisp examples/nil.edn
	./yagolisp examples/bool.edn
	./yagolisp examples/int.edn
	./yagolisp examples/real.edn
	./yagolisp examples/string.edn
	./yagolisp examples/keyword.edn
	./yagolisp examples/symbol.edn

.PHONY: test build clean bench run

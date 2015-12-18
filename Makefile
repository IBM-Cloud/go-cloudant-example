all: test

clean:
	rm -f go-cloudant

install: prepare
	godep go install

prepare:
	go get github.com/tools/godep

build: prepare
	godep go build
	bower install

test: prepare build
	echo "no tests"

.PHONY: install prepare build test

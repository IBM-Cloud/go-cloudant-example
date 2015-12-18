all: test

clean:
	rm -f go-cloudant

install: prepare
	godep go install

prepare:
	go get github.com/tools/godep

build: prepare
	bower install
	godep go build

test: prepare build
	echo "no tests"

.PHONY: install prepare build test

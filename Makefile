all: test

clean:
	rm -f go-cloudant

install: prepare
	godep go install

prepare:
	bower install --config.interactive=false --allow-root
	pwd
	ls -lah
	ls -lah public/vendor
	go get github.com/tools/godep

build: prepare
	godep go build

test: prepare build
	echo "no tests"

.PHONY: install prepare build test

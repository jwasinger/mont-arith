.PHONY: build

build:
	./arith_generator/arith_generator
	bash -c "gofmt -w ."

test:
	go test -v -run=.

benchmark:
	go test -bench=.

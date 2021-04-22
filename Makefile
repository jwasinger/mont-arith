.PHONY: build

build:
	./arith_generator/arith_generator

test:
	go test -v -run=.

benchmark:
	go test -bench=.

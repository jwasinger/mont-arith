.PHONY: build

build:
	./arith_generator/arith_generator
	bash -c "gofmt -w ."

test:
	go test -run=.

benchmark:
	go test -bench=.

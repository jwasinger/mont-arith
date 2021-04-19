.PHONY: build

build:
	rm -rf arith
	mkdir arith
	cp src/* arith/
	./arith_generator/arith_generator
	bash -c "gofmt -w arith/*"

test:
	go test -v -run=.

benchmark:
	go test -bench=.

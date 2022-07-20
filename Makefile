.PHONY: build

build:
	cd arith_generator
	go build
	cd ..
	./arith_generator/arith_generator
	bash -c "gofmt -w ."

test:
	go test -run=.

benchmark:
	go test -bench=.

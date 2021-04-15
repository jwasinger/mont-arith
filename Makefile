.PHONY: build

build:
	rm -rf build
	mkdir build 
	cp src/* build/
	./evmmax-arith-generator
	bash -c "gofmt -w build/*"

test:
	cd build && go test -v -run=MulModMont

benchmark:
	cd build && go test -bench=MulModMont -v

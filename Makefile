.PHONY: build

build:
	rm -rf build
	mkdir build 
	./evmmax-arith-generator
	bash -c "gofmt -w build/*"

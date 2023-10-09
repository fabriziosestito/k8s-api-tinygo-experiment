policy.wasm: $(SOURCE_FILES) go.mod go.sum
	tinygo build -o policy.wasm -target=wasi -gc leaking -no-debug -opt 0 -stack-size 1GB .

.PHONY: clean
clean:
	rm -f policy.wasm

.PHONY: test
test:
	tinygo test -v -target=wasi -no-debug  -opt 0 .


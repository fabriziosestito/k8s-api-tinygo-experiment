policy.wasm: $(SOURCE_FILES) go.mod go.sum
	# docker run \
	# 	--pull=always \
	# 	--rm \
	# 	-e GOFLAGS="-buildvcs=false" \
	# 	-v ${PWD}:/src \
	# 	-w /src ghcr.io/tinygo-org/tinygo-dev:latest \
	tinygo build -o policy.wasm -target=wasi -no-debug -opt 0 .

.PHONY: clean
clean:
	rm -f policy.wasm

.PHONY: test
test:
	/usr/local/tinygo/bin/tinygo test -v  -target=wasi -no-debug .


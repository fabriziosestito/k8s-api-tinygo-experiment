policy.wasm: $(SOURCE_FILES) go.mod go.sum
	docker run \
		--rm \
		-e GOFLAGS="-buildvcs=false" \
		-v ${PWD}:/src \
		-w /src ghcr.io/tinygo-org/tinygo-dev:latest \
		tinygo build -o policy.wasm -target=wasi -no-debug .

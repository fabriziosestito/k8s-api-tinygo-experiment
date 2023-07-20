Experiment compiling k8s.io api with tinygo.

The compilation fails with the following error:

```
docker run \
        --rm \
        -e GOFLAGS="-buildvcs=false" \
        -v /home/lain/wrk/k8s-api-stripped:/src \
        -w /src ghcr.io/tinygo-org/tinygo-dev:latest \
        tinygo build -o policy.wasm -target=wasi -no-debug .
tinygo:wasm-ld: error: lto.tmp: undefined symbol: __multi3
tinygo:wasm-ld: error: lto.tmp: undefined symbol: __multi3
tinygo:wasm-ld: error: lto.tmp: undefined symbol: __multi3
tinygo:wasm-ld: error: lto.tmp: undefined symbol: __multi3
tinygo:wasm-ld: error: lto.tmp: undefined symbol: __multi3
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.unsafe_New
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.typedmemmove
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.mapassign
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.typedslicecopy
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.unsafe_NewArray
tinygo:wasm-ld: error: /tmp/tinygo2762334243/main.o: undefined symbol: reflect.makemap
failed to run tool: wasm-ld
error: failed to link /tmp/tinygo2762334243/main: exit status 1
make: *** [Makefile:2: policy.wasm] Error 1
```

NOTE:

klog has been patched to remove any reference go `os/user` package.
This could be patched upstream by conditionally compile: https://github.com/kubernetes/klog/blob/6632ba5cc9a5331c853a2c6dbc0fbb125e6e7b94/klog_file_others.go#L7
The same approach has been applied to https://github.com/kubernetes/klog/blob/6632ba5cc9a5331c853a2c6dbc0fbb125e6e7b94/klog_file_windows.go

k8s.io/api has been patched to remove `net.ResolveUDPAddr` and `net.ListenUDP` calls, that cause the following error:

```
vendor/k8s.io/utils/net/port.go:116:20: undefined: net.ResolveUDPAddr
vendor/k8s.io/utils/net/port.go:120:20: undefined: net.ListenUDP
```

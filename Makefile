.PHONY: all clean run test

wasms := $(patsubst %.go,%.wasm,$(wildcard wasm/test*/main.go))

all: wasm/main wasm/wasm_exec.js $(wasms)

clean:
	rm -f wasm/main
	rm -f $(wasms)

run: all
	wasm/main

wasm/main: ./*.go wasm/main.go
	go build -o wasm/main wasm/main.go
	
wasm/wasm_exec.js:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .


wasm/test%.wasm: wasm/test%.go
	GOOS=js GOARCH=wasm go build -o $@ $(patsubst %.wasm,%.go,$@)



package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// this application opens a simple file server
func main() {
	port := 9090
	if len(os.Args) > 1 {
		if p, err := strconv.Atoi(os.Args[1]); err == nil {
			port = p
		}
	}
	fmt.Println("FileServer will open on port", port)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if !strings.HasSuffix(path, "/wasm") {
		path += "/wasm"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(path))); err != nil {
		fmt.Printf("Error opening web server on port %d: %s\n", port, err.Error())
		os.Exit(1)
	}
}

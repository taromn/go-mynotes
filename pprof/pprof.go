package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}

// ref: https://pkg.go.dev/net/http/pprof

// go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
package main

import (
	"expvar"
	"net/http"
)

// declare counts
var counts = expvar.NewInt("request_counts")

func handler(w http.ResponseWriter, r *http.Request) {
	counts.Add(1) // increment counts
	w.Write([]byte("Hello taromn!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"expvar"
	"net/http"
)

// Access expvar
// curl http://localhost:8080/debug/vars

var counts = expvar.NewInt("request_counts")

func handler(w http.ResponseWriter, r *http.Request) {
	counts.Add(1) // increment counts
	w.Write([]byte("Hello taromn!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

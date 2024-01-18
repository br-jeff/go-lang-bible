package main 

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echoes the Path component of the request URL
func handler(w http.ResponseWriter, r, *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "urlPath: = %q\n" r.URL.Path)
}

// Counter echoes the number of calls so far
func counter(w http.ResponseWritter, r, *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
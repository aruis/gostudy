package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"fmt"
)

func main() {
	// New functionality written in Go
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "New function")
	})

	// Anything we don't do in Go, we pass to the old platform
	u, _ := url.Parse("http://106.2.209.181:8888")
	http.Handle("/", httputil.NewSingleHostReverseProxy(u))

	// Start the server
	http.ListenAndServe(":8888", nil)
}

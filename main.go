package main

import (
	"fmt"
	"net/http"
)

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "post\n")
}

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}
	srv.ListenAndServe()
}

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>I made it!)</h1>"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

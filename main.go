//go:generate go run gen/main.go

package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// here we use our embedded asset
		w.Write(image)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

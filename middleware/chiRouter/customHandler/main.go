package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/test", Handler(customHandler))

	addr := ":3333"
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

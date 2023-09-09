// Versions
// ========
// This example demonstrates the use of the render subpackage, with
// a quick concept for how to support multiple api versions.
package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/_examples/versions/data"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API version 3.
	r.Route("/v3", func(r chi.Router) {
		r.Use(apiVersionCtx("v3"))
		r.Mount("/articles", articleRouter())
	})

	// API version 2.
	r.Route("/v2", func(r chi.Router) {
		r.Use(apiVersionCtx("v2"))
		r.Mount("/articles", articleRouter())
	})

	// API version 1.
	r.Route("/v1", func(r chi.Router) {
		r.Use(randomErrorMiddleware) // Simulate random error, ie. version 1 is buggy.
		r.Use(apiVersionCtx("v1"))
		r.Mount("/articles", articleRouter())
	})

	addr := ":3333"
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func randomErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())

		// One in three chance of random error.
		if rand.Int31n(3) == 0 {
			errors := []error{data.ErrUnauthorized, data.ErrForbidden, data.ErrNotFound}
			render.Respond(w, r, errors[rand.Intn(len(errors))])
			return
		}
		next.ServeHTTP(w, r)
	})
}

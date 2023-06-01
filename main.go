package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my url-shortener</h1>")
}

func urlContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlId := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), "urlId", &urlId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("urlId").(*string)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
	} else {
		http.Redirect(w, r, "https://music.youtube.com", 301)
	}
}

func main() {
	r := chi.NewRouter()

	// TODO: expand middleware as needed
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Register the routes
	r.Get("/", homeHandler)

	r.Route("/api/{id}", func(r chi.Router) {
		r.Use(urlContext)
		r.Get("/", getUrl)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}

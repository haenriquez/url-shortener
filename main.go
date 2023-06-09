package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/haenriquez/url-shortener/controllers"
	"github.com/haenriquez/url-shortener/views"
	"net/http"
	"path/filepath"
)

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

func postUrl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form submission", http.StatusBadRequest)
		return
	}

	fmt.Println("url: ", r.FormValue("url"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed successfully"))
}

func main() {
	r := chi.NewRouter()

	// TODO: expand middleware as needed
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	// r.Use(middleware.Timeout(60 * time.Second))

	// Parse templates before starting the server
	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))),
	)
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))),
	)
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))),
	)

	r.Route("/api/{id}", func(r chi.Router) {
		r.Use(urlContext)
		r.Get("/", getUrl)
	})

	r.Route("/api/shorten", func(r chi.Router) {
		r.Post("/", postUrl)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

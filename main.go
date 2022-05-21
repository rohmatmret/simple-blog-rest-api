package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/simple-blog/post/handler"
	"github.com/simple-blog/post/repository"
)

func main() {
	port := ":3000"
	runtime.GOMAXPROCS(2)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Timeout(60 * time.Second))
	db := Connection()

	posthandler := handler.NewPostHandler(r, repository.NewPostRepository(db))
	r.Route("/", func(r chi.Router) {
		r.Get("/", posthandler.FindAll)
		r.Post("/", posthandler.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", posthandler.FindById)
			r.Put("/", posthandler.Update)
			r.Delete("/", posthandler.Deleted)
		})
	})
	log.Printf("Starting up on http://localhost:%s", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Panic(err)
	}
}

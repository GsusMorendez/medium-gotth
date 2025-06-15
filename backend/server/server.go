package server

import (
	"fmt"
	"log"
	"net/http"
	"version1-medium-gotth/backend/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(port int) {
	painter := NewPainter()
	data := repository.NewDataManager()
	handler := NewHandler(painter, data)
	server := NewServer(handler)

	log.Printf("Server listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), server); err != nil {
		log.Fatal(err)
	}
}

func NewServer(handler *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Handle("/frontend/static/*", http.StripPrefix("/frontend/static/", http.FileServer(http.Dir("frontend/static"))))
	setupRoutes(r, handler)
	return r
}

func setupRoutes(r chi.Router, handler *Handler) {
	r.Get("/hello", handler.Hello)
	r.Get("/main", handler.MainView)
}

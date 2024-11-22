package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sean-d/tentbnb/pkg/config"
	"github.com/sean-d/tentbnb/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/parry", handlers.Repo.Parry)
	mux.Get("/simon", handlers.Repo.Simon)

	// handling serving the static files
	// create a Handler, fileServer using the contents of the filesystem location supplied
	// Handle adds a route for all things in /static/ and uses a Handler that serves files from
	// a filesystem dir as a replacement for /static/
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}

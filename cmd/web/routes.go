package main

import (
	"gethub.com/atobiason/bookings/pkg/config"
	"gethub.com/atobiason/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	/*
		mux := pat.New()
		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
		return mux
	*/
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

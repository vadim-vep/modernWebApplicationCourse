package main

import (

	"github.com/go-chi/chi"
	"modernWebAppCourse/pkg/config"
	"modernWebAppCourse/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/", http.HandlerFunc(handlers.Repo.About))
	return mux
}

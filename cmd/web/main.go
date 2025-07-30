package main

import (
	"fmt"
	"log"
	"modernWebAppCourse/pkg/config"
	"modernWebAppCourse/pkg/handlers"
	"modernWebAppCourse/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Server started on port: %s\n", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)


}

type TemplateData struct {
	name string
	Age int
	Address string
}


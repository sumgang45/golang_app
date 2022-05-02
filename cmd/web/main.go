package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sumgang45/LearningGo/webapp/pkg/config"
	"github.com/sumgang45/LearningGo/webapp/pkg/handlers"
	"github.com/sumgang45/LearningGo/webapp/pkg/render"
)

const portNumber = ":8080"

//this is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

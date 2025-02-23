package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mariaptv/WebAppGo/pkg/config"
	"github.com/mariaptv/WebAppGo/pkg/handlers"
	"github.com/mariaptv/WebAppGo/pkg/render"
)

const portNumber = ":8080"

// main is the main function
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

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: route(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

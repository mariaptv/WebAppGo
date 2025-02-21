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
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

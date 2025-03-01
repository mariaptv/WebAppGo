package config

import (
	"log"
	"text/template"
)

//Dosen´t import anything outside of the application

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog     *log.Logger
	InProduction bool
}
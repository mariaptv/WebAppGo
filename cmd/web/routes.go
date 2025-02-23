package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mariaptv/WebAppGo/pkg/config"
	"github.com/mariaptv/WebAppGo/pkg/handlers"
)
func route (app *config.AppConfig) http.Handler{
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
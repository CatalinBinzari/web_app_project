package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portnumber = ":8080"

func main() {
	// fmt.Println("hi")

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create tmpl cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// pass app config to handlers pkg
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// gives to render pkg access to app.config
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

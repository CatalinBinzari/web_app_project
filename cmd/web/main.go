package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portnumber = ":8080"

// Global variables used in main pkg
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change it to true when in production
	app.InProduction = false

	// sessions by default are stored in memory, can be used databases also
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // in prod to be set to true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create tmpl cache")
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

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

package main

import (
	"encoding/gob"
	"log"
	"myapp/internal/config"
	"myapp/internal/handlers"
	"myapp/internal/helpers"
	"myapp/internal/models"
	"myapp/internal/render"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portnumber = ":8080"

// Global variables used in main pkg
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {

	// what we store in the session
	gob.Register(models.Reservation{})

	// change it to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	app.ErrorLog = errorLog

	// sessions by default are stored in memory, can be used databases also
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // in prod to be set to true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("could not create tmpl cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	// pass app config to handlers pkg
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// gives to render pkg access to app.config
	render.NewTemplates(&app)

	helpers.NewHelpers(&app)

	return nil
}

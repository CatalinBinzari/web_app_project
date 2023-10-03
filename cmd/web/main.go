package main

import (
	"encoding/gob"
	"log"
	"myapp/internal/config"
	"myapp/internal/driver"
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
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run() (*driver.DB, error) {

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

	// connect to Db
	log.Println("Connection to database initialized...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=postgres user=postgres password=123")
	if err != nil {
		log.Fatal("cannot connect to db, dye: ", db)
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("could not create tmpl cache")
		return nil, err
	}
	log.Println("Connected to database with success")

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	// pass app config to handlers pkg
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	// gives to render pkg access to app.config
	render.NewTemplates(&app)

	helpers.NewHelpers(&app)

	return db, nil
}

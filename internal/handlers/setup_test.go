package handlers

import (
	"fmt"
	"html/template"
	"myapp/internal/config"
	"net/http"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

// func getRoutes() http.Handler {

// 	// what we store in the session
// 	gob.Register(models.Reservation{})

// 	// change it to true when in production
// 	app.InProduction = false

// 	// sessions by default are stored in memory, can be used databases also
// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction // in prod to be set to true

// 	app.Session = session

// 	tc, err := CreateTestTemplateCache()
// 	if err != nil {
// 		log.Println("could not create tmpl cache")
// 	}

// 	app.TemplateCache = tc
// 	// app.UseCache = app.InProduction
// 	app.UseCache = true

// 	// pass app config to handlers pkg
// 	repo := NewRepo(&app)
// 	NewHandlers(repo)

// 	// gives to render pkg access to app.config
// 	render.NewTemplates(&app)
// 	mux := chi.NewRouter()

// 	mux.Use(middleware.Recoverer)
// 	mux.Use(NoSurf)
// 	mux.Use(SessionLoad)

// 	mux.Get("/", Repo.Home)
// 	mux.Get("/about", Repo.About)
// 	mux.Get("/generals-quarters", Repo.Generals)
// 	mux.Get("/majors-suite", Repo.Majors)

// 	mux.Get("/search-availability", Repo.Availability)
// 	mux.Post("/search-availability", Repo.PostAvailability)
// 	mux.Get("/search-availability-json", Repo.AvailabilityJSON)
// 	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

// 	mux.Get("/contact", Repo.Contact)

// 	mux.Get("/make-reservation", Repo.Reservation)
// 	mux.Post("/make-reservation", Repo.PostReservation)
// 	mux.Get("/reservation-summary", Repo.ReservationSummary)

// 	fileServer := http.FileServer(http.Dir("./static/"))
// 	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

// 	return mux
// }

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{} // the same as make(map[string]*template.Template)

	// get all files *.page.tmpl
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all files
	for _, page := range pages {
		name := filepath.Base(page) // pages is full path, we tak the filename
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

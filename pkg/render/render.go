package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	// parse tmpl
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Err", err)
// 		return
// 	}
// }

var tc = make(map[string]*template.Template) // template cache

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// see if we have it in the cache already, if inMap is true then template found, else false
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating tmpl and adding to cache")
		err = CreateTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//
	tc[t] = tmpl

	return nil
}

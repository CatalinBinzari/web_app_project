package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portnumber = ":8080"

func main() {
	// fmt.Println("hi")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port %s", portnumber)
	_ = http.ListenAndServe(portnumber, nil)

}

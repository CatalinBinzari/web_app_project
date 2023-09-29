package helpers

import (
	"fmt"
	"myapp/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers sets up ap config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError writes to w the client error, and saves it to app.InfoLog
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError prints error message and stack, and saves it to app.ErrorLog
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

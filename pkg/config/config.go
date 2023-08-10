package config

import (
	"html/template"
	"log"
)

// config is imported by other parts of the app, but it does not import anything

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}

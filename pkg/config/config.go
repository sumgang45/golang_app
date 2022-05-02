package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//holds application config
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InfoLog        *log.Logger
	InProduction   bool
	SessionManager *scs.SessionManager
}

package handlers

import (
	"net/http"

	"github.com/sumgang45/LearningGo/webapp/pkg/config"
	"github.com/sumgang45/LearningGo/webapp/pkg/models"
	"github.com/sumgang45/LearningGo/webapp/pkg/render"
)

//repository used by the handlers
var Repo *Repository

//Repository type
type Repository struct {
	App *config.AppConfig
}

//Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//this is an about page and it also receives information from the addvalues private func
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

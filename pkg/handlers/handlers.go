package handlers

import (
	"net/http"

	"github.com/sumgang45/LearningGo/webapp/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

//this is an about page and it also receives information from the addvalues private func
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

//renders templares using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

	parseTemplate, _ := template.ParseFiles(".\\templates\\" + tmpl)

	err = parseTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

//Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	//map: a variable to look up objects very quickly. It is a cache to look up all templates and then render the appropriate per req
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(".\\templates\\*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("Page is currently", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(".\\templates\\*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(".\\templates\\*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

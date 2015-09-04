package controllers

import (
	"github.com/gpavlidi/gowebsite/helpers"
	"html/template"
	"net/http"
)

var staticTemplates map[string]*template.Template

func init() {
	templates := helpers.GetSharedTemplates()
	staticTemplates = map[string]*template.Template{
		"index":   template.Must(template.ParseFiles(append(templates, "views/static/index.html")...)),
		"contact": template.Must(template.ParseFiles(append(templates, "views/static/contact.html")...)),
	}

}

func StaticIndex(w http.ResponseWriter, r *http.Request) {
	type Page struct{ Title string }
	p := &Page{Title: "Index"}
	err := staticTemplates["index"].ExecuteTemplate(w, "base", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StaticContact(w http.ResponseWriter, r *http.Request) {
	type Page struct{ Title string }
	p := &Page{Title: "Contact"}
	err := staticTemplates["contact"].ExecuteTemplate(w, "base", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

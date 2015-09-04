package controllers

import (
	"github.com/gpavlidi/gowebsite/helpers"
	"html/template"
	"net/http"
)

var userTemplates map[string]*template.Template

func init() {
	templates := helpers.GetSharedTemplates()
	userTemplates = map[string]*template.Template{
		"index": template.Must(template.ParseFiles(append(templates, "views/user/index.html")...)),
	}

}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	type Page struct{ Title string }
	p := &Page{Title: "Users"}
	err := userTemplates["index"].ExecuteTemplate(w, "base", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

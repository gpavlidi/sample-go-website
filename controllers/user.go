package controllers

import (
	"html/template"
	"net/http"

	"github.com/gpavlidi/go-website/helpers"
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

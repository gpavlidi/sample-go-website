package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gpavlidi/go-website/helpers"
	"github.com/gpavlidi/go-website/models"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var userTemplates map[string]*template.Template

func init() {
	templates := helpers.GetSharedTemplates()
	userTemplates = map[string]*template.Template{
		"index": template.Must(template.ParseFiles(append(templates, "views/user/index.html")...)),
		"show":  template.Must(template.ParseFiles(append(templates, "views/user/show.html")...)),
		"edit":  template.Must(template.ParseFiles(append(templates, "views/user/edit.html")...)),
		"new":   template.Must(template.ParseFiles(append(templates, "views/user/new.html")...)),
	}

}

/*
	View Endpoints
*/
func UserIndex(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*sql.DB)

	users, err := models.GetAllUsers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = userTemplates["index"].ExecuteTemplate(w, "base", users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserNew(w http.ResponseWriter, r *http.Request) {
	err := userTemplates["new"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	db := context.Get(r, "db").(*sql.DB)
	user, err := models.GetUserById(db, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = userTemplates["show"].ExecuteTemplate(w, "base", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	db := context.Get(r, "db").(*sql.DB)
	user, err := models.GetUserById(db, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = userTemplates["edit"].ExecuteTemplate(w, "base", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
	API Endpoints:
*/
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := &models.User{}
	user.Id = vars["id"]
	user.FirstName = r.FormValue("user[first_name]")
	user.LastName = r.FormValue("user[last_name]")

	db := context.Get(r, "db").(*sql.DB)

	err := models.NewOrUpdateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", 301)
}

func UserPost(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	user.FirstName = r.FormValue("user[first_name]")
	user.LastName = r.FormValue("user[last_name]")

	db := context.Get(r, "db").(*sql.DB)

	err := models.NewOrUpdateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", 301)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	db := context.Get(r, "db").(*sql.DB)

	err := models.DeleteUserById(db, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", 301)
}

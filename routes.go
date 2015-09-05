package main

import (
	"net/http"

	"github.com/gpavlidi/go-website/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"StaticIndex", "GET", "/", controllers.StaticIndex},
	Route{"StaticContact", "GET", "/contact", controllers.StaticContact},

	Route{"UserIndex", "GET", "/users", controllers.UserIndex},
	Route{"UserShow", "GET", "/users/{id:[0-9]+}", controllers.UserShow},
	Route{"UserEdit", "GET", "/users/{id:[0-9]+}/edit", controllers.UserEdit},
	Route{"UserNew", "GET", "/users/new", controllers.UserNew},
	//API endpoints
	Route{"UserPost", "POST", "/users", controllers.UserPost},
	Route{"UserUpdate", "PATCH", "/users/{id:[0-9]+}", controllers.UserUpdate},
	Route{"UserDelete", "DELETE", "/users/{id:[0-9]+}", controllers.UserDelete},
}

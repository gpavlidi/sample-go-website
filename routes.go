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
}

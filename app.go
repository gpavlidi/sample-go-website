package main

import (
	"database/sql"
	//_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/julienschmidt/httprouter"
	mw "github.com/gpavlidi/go-website/middlewares"
)

type AppConfig struct {
	Addr string
	Name string
}

type App struct {
	Handler http.Handler /**httprouter.Router*/ /**mux.Router*/ /*gorilla, *http.ServeMux is the golang native*/
	Log     *log.Logger
	Config  *AppConfig
	Db      *sql.DB
}

func NewApp(cfg *AppConfig) *App {
	log := log.New(os.Stdout, cfg.Name, log.LstdFlags)

	router := newRouter(routes)

	//Db := NewDb("postgres", "postgres://user:pass@localhost:5432/goboostrap?sslmode=disable")
	//Db := NewDb("sqlite3", ":memory:")
	Db := NewDb("sqlite3", "./site.db")

	handler := mw.UseOn(router, mw.SetDB(Db), handlers.HTTPMethodOverrideHandler, mw.LogRequest(os.Stdout), handlers.CompressHandler, mw.SetHeader("test1", "boom"), context.ClearHandler)

	app := &App{handler, log, cfg, Db}

	return app
}

func NewDb(driver, dsn string) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Pass our own mux to ListenAndServe
// https://gowalker.org/net/http#ListenAndServe
func (app *App) ListenAndServe() {
	http.ListenAndServe(app.Config.Addr, app.Handler)
	app.Db.Close()
}

func newRouter(routes Routes) *mux.Router /**httprouter.Router *mux.Router*/ {
	// switched back to gorilla mux because it keeps handlers compatible with http.HandlerFunc
	// because it carries the params through the mutex
	// httprouter adds a 3rd parameter to HandlerFunc for the params, breaking the interface
	//

	// native golang mux
	/*mux := http.NewServeMux()
	for _, rt := range routes {
		mux.HandleFunc(rt.Pattern, rt.HandlerFunc)
	}*/
	//gorilla mux, offers more stuff like per Method routes
	router := mux.NewRouter().StrictSlash(true)
	for _, rt := range routes {
		router.
			Methods(rt.Method).
			Path(rt.Pattern).
			Name(rt.Name).
			Handler(rt.HandlerFunc)
	}
	// Path of static files must be last!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	/*router := httprouter.New()
	for _, rt := range routes {
		//router.Handler(rt.Method, rt.Pattern, middlewares.SetHeader(rt.HandlerFunc))
		router.HandlerFunc(rt.Method, rt.Pattern, rt.HandlerFunc)
	}
	// hack to get httprouter to server public from root
	// https://github.com/julienschmidt/httprouter/issues/4#issuecomment-41549684
	router.NotFound = http.FileServer(http.Dir("public")) */

	return router
}

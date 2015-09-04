package middlewares

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func SetHeader(headerKey, headerVal string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set(headerKey, headerVal)

			next.ServeHTTP(res, req)
		})
	}
}

// set db pointer on the request's context so that it's available everywhere
func SetDB(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			context.Set(req, "db", db)

			next.ServeHTTP(res, req)
		})
	}
}

// wrap gorilla's logging middleware to follow the calling convention standards
func LogRequest(logFile *os.File) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(logFile, next)
	}
}

// app is the final handler usually the mux
// handlers are applied from left to right
// e.g. first handler is applied first and exits last
func UseOn(app http.Handler, handlers ...func(next http.Handler) http.Handler) http.Handler {
	handler := app
	for ind, _ := range handlers {
		hlr := handlers[len(handlers)-ind-1]
		handler = hlr(handler)
	}
	return handler
}

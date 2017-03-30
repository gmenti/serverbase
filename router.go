package main

import (
	"net/http"

	"github.com/gmenti/serverbase/controllers/userController"
	"github.com/gorilla/mux"
	"time"
	"log"
)

type Route struct {
	Name		string
	Method  	string
	Pattern		string
	HandlerFunc 	http.HandlerFunc
}

var routes = []Route{
	{
		"UserIndex",
		"GET",
		"/users",
		userController.Index,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		handler := http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
			startedTime := time.Now()

			writer.Header().Add("Content-Type", "application/json")
			route.HandlerFunc.ServeHTTP(writer, request)

			log.Printf(
				"%s\t%s\t%s\t%s",
				request.Method,
				request.RequestURI,
				route.Name,
				time.Since(startedTime),
			)
		})

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handler)

	}

	return router
}
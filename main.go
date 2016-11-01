package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/tamizhvendan/golang-web-boilerplate/handlers"
	"github.com/tamizhvendan/golang-web-boilerplate/middlewares"
	"github.com/urfave/negroni"
)

func main() {
	router := httprouter.New()
	indexHandler := &handlers.IndexHandler{Greeting: "Hello"}
	indexHandler.AddRoutes(router)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewares.RequestId))
	n.Use(negroni.HandlerFunc(middlewares.AuthScopes))
	n.UseHandler(context.ClearHandler(router))
	n.Run(":3000")
}

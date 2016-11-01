package handlers

import (
	"net/http"

	"github.com/deckarep/golang-set"
	"github.com/julienschmidt/httprouter"
	. "github.com/tamizhvendan/golang-web-boilerplate/infrastructure"
)

type IndexHandler struct {
	Greeting string // Dependency!
}

func (h *IndexHandler) AddRoutes(r *httprouter.Router) {
	r.GET("/greet/:name", AppErrorHandler(h.Greet))
	securedGreetScopes := mapset.NewSet("greet.other")
	r.GET("/secured_greet/:name", AppErrorHandler(HasScopes(securedGreetScopes, h.Greet)))
}

func (h *IndexHandler) Greet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *AppError {
	name := ps.ByName("name")
	if name == "foo" {
		return &AppError{
			StatusCode: http.StatusBadRequest,
			Message:    "name should not be foo",
		}
	}
	var greeting struct {
		Message string
	}
	greeting.Message = h.Greeting + " " + name
	WriteJson(http.StatusOK, greeting, w)
	return nil
}

package infrastructure

import (
	"net/http"

	"github.com/deckarep/golang-set"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/tamizhvendan/golang-web-boilerplate/middlewares"
)

func HasScopes(requiredScopes mapset.Set, fn handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) *AppError {
		scopes, ok := context.GetOk(r, middlewares.AuthScopesKey)
		actualScopes, ok2 := scopes.(mapset.Set)
		if !ok || !ok2 {
			return &AppError{
				Message:    "Error while getting scope",
				StatusCode: http.StatusInternalServerError,
			}
		}
		if !actualScopes.IsSuperset(requiredScopes) {
			return &AppError{
				Message:    "Access denied",
				StatusCode: http.StatusForbidden,
			}
		}
		return fn(w, r, p)
	}
}

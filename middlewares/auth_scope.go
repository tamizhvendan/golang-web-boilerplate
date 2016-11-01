package middlewares

import (
	"net/http"
	"strings"

	"github.com/deckarep/golang-set"
	"github.com/gorilla/context"
)

const AuthScopesKey = "authScopes"

func AuthScopes(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authToken := r.Header.Get("Authorization")
	scopes := mapset.NewSet()
	if strings.HasPrefix(authToken, "Bearer ") {
		token := strings.Replace(authToken, "Bearer ", "", -1)
		scopes = getScopes(token)
	}
	context.Set(r, AuthScopesKey, scopes)
	next(rw, r)
}

func getScopes(token string) mapset.Set {
	// Use JWT or talk to your Identity server to get the scopes from token
	scopes := mapset.NewSet()
	if token == "foobar" {
		scopes.Add("greet.me")
		scopes.Add("greet.other")
	}
	return scopes
}

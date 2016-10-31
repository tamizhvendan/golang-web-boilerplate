package middlewares

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/satori/go.uuid"
)

func RequestId(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestId := uuid.NewV4()
	context.Set(r, "requestId", requestId)
	next(rw, r)
}

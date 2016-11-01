package infrastructure

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/tamizhvendan/golang-web-boilerplate/middlewares"
)

type AppError struct {
	Message    string
	StatusCode int
}

type handlerFunc func(http.ResponseWriter, *http.Request, httprouter.Params) *AppError

func AppErrorHandler(fn handlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		appErr := fn(w, r, p)
		requestId := context.Get(r, middlewares.RequestIdKey)
		if appErr != nil {
			writeAppError(appErr, requestId, w, r)
		}
	}
}

func writeAppError(appErr *AppError, requestId interface{}, w http.ResponseWriter, r *http.Request) {
	log.WithField("requestId", requestId).Error(appErr.Message)
	var httpError struct {
		*AppError
		RequestId interface{}
	}
	httpError.AppError = appErr
	httpError.RequestId = requestId
	WriteJson(appErr.StatusCode, httpError, w)
}

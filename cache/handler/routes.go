package handler

import (
	"github.com/jpastorm/redis-cache/cache/usecase"
	"net/http"
)

const (
	adminRoutesPrefix = "/api/v1/cache"
)

// NewRouter returns a router to handle model.EndpointJob from a client
func NewRouter(app *http.ServeMux, useCase usecase.Usecase) {
	handle := NewHandler(useCase)
	adminRoutes(app, handle)
}

func adminRoutes(app *http.ServeMux, handle handler) {
	app.HandleFunc(adminRoutesPrefix+"/set", handle.Set)
	app.HandleFunc(adminRoutesPrefix+"/get", handle.Get)
}
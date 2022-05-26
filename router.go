package router

import (
	"context"
	"fmt"
	application "github.com/debugger84/modulus-application"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
)

type Router struct {
	router *httprouter.Router
	port   int
	logger application.Logger
}

func NewRouter(config *ModuleConfig, logger application.Logger) *Router {
	r := &httprouter.Router{
		RedirectTrailingSlash:  config.RedirectTrailingSlash,
		RedirectFixedPath:      config.RedirectFixedPath,
		HandleMethodNotAllowed: config.HandleMethodNotAllowed,
		HandleOPTIONS:          config.HandleOPTIONS,
		GlobalOPTIONS:          config.GlobalOPTIONS,
		NotFound:               config.NotFound,
		MethodNotAllowed:       config.MethodNotAllowed,
		PanicHandler:           config.PanicHandler,
	}
	router := &Router{router: r, logger: logger}
	router.AddRoutes(config.Routes.GetRoutesInfo())
	router.port = config.Port
	return router
}

func (r *Router) AddRoutes(routes []application.RouteInfo) {
	for _, info := range routes {
		r.router.Handler(info.Method(), info.Path(), info.Handler())
	}
}

func (r *Router) Run() error {
	r.logger.Info(context.Background(), fmt.Sprintf("Listen to the port: %d", r.port), "port", r.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.router)
}

func (r *Router) RouteParams(request *http.Request) url.Values {
	result := make(url.Values)
	params := httprouter.ParamsFromContext(request.Context())
	for _, param := range params {
		result.Add(param.Key, param.Value)
	}
	return result
}

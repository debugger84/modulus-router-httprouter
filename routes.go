package router

import (
	application "github.com/debugger84/modulus-application"
	"net/http"
)

type Routes struct {
	routes map[string]application.RouteInfo
}

func NewRoutes() *Routes {
	return &Routes{routes: make(map[string]application.RouteInfo)}
}

func (r *Routes) Get(name string, path string, handler http.HandlerFunc) {
	r.routes[name] = *application.NewRouteInfo(
		http.MethodGet,
		path,
		handler,
	)
}

func (r *Routes) Post(name string, path string, handler http.HandlerFunc) {
	r.routes[name] = *application.NewRouteInfo(
		http.MethodPost,
		path,
		handler,
	)
}

func (r *Routes) Delete(name string, path string, handler http.HandlerFunc) {
	r.routes[name] = *application.NewRouteInfo(
		http.MethodDelete,
		path,
		handler,
	)
}

func (r *Routes) Put(name string, path string, handler http.HandlerFunc) {
	r.routes[name] = *application.NewRouteInfo(
		http.MethodPut,
		path,
		handler,
	)
}

func (r *Routes) Options(name string, path string, handler http.HandlerFunc) {
	r.routes[name] = *application.NewRouteInfo(
		http.MethodOptions,
		path,
		handler,
	)
}

func (r *Routes) AddFromRoutes(routes *Routes) {
	for name, info := range routes.routes {
		r.routes[name] = info
	}
}

func (r *Routes) GetRoutesInfo() []application.RouteInfo {
	result := make([]application.RouteInfo, 0, len(r.routes))
	for _, info := range r.routes {
		result = append(result, info)
	}

	return result
}

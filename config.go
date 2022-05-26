package router

import (
	application "github.com/debugger84/modulus-application"
	"go.uber.org/dig"
	"net/http"
)

type ModuleConfig struct {
	Routes                 Routes
	Port                   int
	RedirectTrailingSlash  bool
	RedirectFixedPath      bool
	HandleMethodNotAllowed bool
	HandleOPTIONS          bool
	GlobalOPTIONS          http.Handler
	NotFound               http.Handler
	MethodNotAllowed       http.Handler
	PanicHandler           func(http.ResponseWriter, *http.Request, interface{})
	container              *dig.Container
}

func NewModuleConfig() *ModuleConfig {
	return &ModuleConfig{
		Routes:                 *NewRoutes(),
		Port:                   0,
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: false,
		HandleOPTIONS:          false,
		GlobalOPTIONS:          nil,
		NotFound:               nil,
		MethodNotAllowed:       nil,
		PanicHandler:           nil,
	}
}

func (c *ModuleConfig) InitConfig(config application.Config) error {
	if c.Port == 0 {
		c.Port = config.GetEnvAsInt("APP_PORT")
	}
	return nil
}

func (c *ModuleConfig) ProvidedServices() []interface{} {
	return []interface{}{
		NewRouter,
		func() *ModuleConfig {
			return c
		},
		func(router *Router) application.Router {
			return router
		},
	}
}

func (c *ModuleConfig) OnStart() error {
	var router *Router
	err := c.container.Invoke(func(dep *Router) error {
		router = dep
		return nil
	})
	if err != nil {
		return err
	}

	return router.Run()
}

func (c *ModuleConfig) SetContainer(container *dig.Container) {
	c.container = container
}

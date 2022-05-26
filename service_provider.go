package router

import (
	application "github.com/debugger84/modulus-application"
	"go.uber.org/dig"
)

type ServiceProvider struct {
	container *dig.Container
	config    *ModuleConfig
}

func (s *ServiceProvider) InitConfig(config application.Config) error {
	if s.config.Port == 0 {
		s.config.Port = config.GetEnvAsInt("APP_PORT")
	}
	return nil
}

func NewServiceProvider(config *ModuleConfig) *ServiceProvider {
	return &ServiceProvider{config: config}
}

func (s *ServiceProvider) Provide() []interface{} {
	return []interface{}{
		NewRouter,
		func() *ModuleConfig {
			return s.config
		},
		func(router *Router) application.Router {
			return router
		},
	}
}

func (s *ServiceProvider) OnStart() error {
	var router *Router
	err := s.container.Invoke(func(dep *Router) error {
		router = dep
		return nil
	})
	if err != nil {
		return err
	}

	return router.Run()
}

func (s *ServiceProvider) OnClose() error {
	return nil
}

func (s *ServiceProvider) SetContainer(container *dig.Container) {
	s.container = container
}

package context

import (
	"github.com/gambitier/goblog/services"
	"github.com/gambitier/goblog/startup/config"
)

type AppContext struct {
	Services *services.Service
	Config   *config.Config
}

func NewAppContext() (*AppContext, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	service, err := services.NewService()
	if err != nil {
		return nil, err
	}

	return &AppContext{
		Services: service,
		Config:   cfg,
	}, nil
}

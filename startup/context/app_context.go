package context

import (
	"fmt"

	"github.com/gambitier/goblog/api/handlers"
	"github.com/gambitier/goblog/database"
	"github.com/gambitier/goblog/services"
	"github.com/gambitier/goblog/startup/config"
)

type AppContext struct {
	Services *services.Service
	Config   *config.Config
	DB       *database.Database
	Handlers *handlers.Handlers
}

func NewAppContext() (*AppContext, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	db, err := database.NewDatabase(&cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	service, err := services.NewService(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize service: %v", err)
	}

	handlers := handlers.NewHandlers(service)

	return &AppContext{
		Services: service,
		Config:   cfg,
		DB:       db,
		Handlers: handlers,
	}, nil
}

package context

import "github.com/gambitier/goblog/services"

type AppContext struct {
	Services *services.Service
}

func NewAppContext() (*AppContext, error) {
	// db, err := database.NewDatabaseRepository()
	// if err != nil {
	// 	return nil, err
	// }

	service, err := services.NewService()
	if err != nil {
		return nil, err
	}

	return &AppContext{
		Services: service,
		// DatabaseRepo: db,
	}, nil
}

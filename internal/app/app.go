package app

import (
	"log"

	"github.com/joho/godotenv"
)

type App struct {
	Usecases     *Usecases
	Repositories *repositories
	Drivers      *drivers
	Services     *services
}

func New() (*App, error) {
	app := new(App)

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env  file \n%v", err)
	}

	drivers, err := newDrivers()
	if err != nil {
		log.Fatalln(err)
	}

	repos := newRepositories(drivers)
	services := newServices()
	ucs := newUsecases(repos, services)

	app.Drivers = drivers
	app.Repositories = repos
	app.Usecases = ucs
	app.Services = services
	return app, nil
}

func (app App) Close() error {
	var closeErr error

	if err := app.Drivers.MySQL.Close(); err != nil {
		log.Println("Failed to close MySQL DB connection")
		closeErr = err
	}

	return closeErr
}

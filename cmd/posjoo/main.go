package main

import (
	"log"

	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("Failed initiate the app\n%v", err)
	}
	defer app.Close()

	web.Start(app)
}

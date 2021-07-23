package web

import (
	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web/controller"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web/server"
)

func Start(app *app.App) {
	web := server.New()

	authController := controller.NewAuthController(app.Usecases)
	authGroup := web.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
	server.Start(web)
}

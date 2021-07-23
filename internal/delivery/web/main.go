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

	// JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	// JWTMiddleware := middleware.JWT(JWTSecretKey)
	// authenticatedGroup := web.Group("", JWTMiddleware)

	server.Start(web)
}

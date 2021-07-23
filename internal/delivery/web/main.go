package web

import (
	"os"

	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web/controller"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web/middleware"
	"github.com/ardafirdausr/posjoo-server/internal/delivery/web/server"
)

func Start(app *app.App) {
	web := server.New()
	api := web.Group("/api/v1")

	authController := controller.NewAuthController(app.Usecases)
	authGroup := api.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTMiddleware := middleware.JWT(JWTSecretKey)

	userController := controller.NewUserController(app.Usecases)
	userGroup := api.Group("/users", JWTMiddleware)
	userGroup.GET("/:userID", userController.GetUser)
	userGroup.GET("", userController.GetMerchantUsers)
	userGroup.POST("/:userID", userController.CreateUser)
	userGroup.PUT("/:userID", userController.UpdateUser)
	userGroup.DELETE("/:userID", userController.DeleteUser)

	server.Start(web)
}

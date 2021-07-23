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
	userGroup.POST("", userController.CreateUser)
	userGroup.PUT("/:userID", userController.UpdateUser)
	// userGroup.PUT("/:userID/password", userController.UpdateUserPassword)
	// userGroup.PUT("/:userID/photo", userController.UpdateUserPhoto)
	userGroup.DELETE("/:userID", userController.DeleteUser)

	productController := controller.NewProductController(app.Usecases)
	productGroup := api.Group("/products", JWTMiddleware)
	productGroup.GET("/:productID", productController.GetProduct)
	productGroup.GET("", productController.GetMerchantProducts)
	productGroup.POST("", productController.CreateProduct)
	productGroup.PUT("/:productID", productController.UpdateProduct)
	// productGroup.PUT("/:productID/photo", productController.UpdateProductPhoto)
	productGroup.DELETE("/:productID", productController.DeleteProduct)

	server.Start(web)
}

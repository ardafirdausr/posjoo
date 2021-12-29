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
	web.Static("/storage", "storage")

	api := web.Group("/api/v1")

	authController := controller.NewAuthController(app.Usecases)
	authGroup := api.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTMiddleware := middleware.JWT(JWTSecretKey)

	OwnerAccess := middleware.OwnerAccess()

	userController := controller.NewUserController(app.Usecases)
	userGroup := api.Group("/users", JWTMiddleware)
	userGroup.GET("/:userID", userController.GetUser, OwnerAccess)
	userGroup.GET("", userController.GetMerchantUsers, OwnerAccess)
	userGroup.POST("", userController.CreateUser, OwnerAccess)
	userGroup.PUT("/:userID", userController.UpdateUser, OwnerAccess)
	userGroup.PUT("/:userID/password", userController.UpdateUserPassword, OwnerAccess)
	userGroup.PUT("/:userID/photo", userController.UpdateUserPhoto, OwnerAccess)
	userGroup.DELETE("/:userID", userController.DeleteUser, OwnerAccess)

	productController := controller.NewProductController(app.Usecases)
	productGroup := api.Group("/products", JWTMiddleware)
	productGroup.GET("/:productID", productController.GetProduct)
	productGroup.GET("", productController.GetMerchantProducts)
	productGroup.POST("", productController.CreateProduct, OwnerAccess)
	productGroup.PUT("/:productID", productController.UpdateProduct, OwnerAccess)
	productGroup.PUT("/:productID/photo", productController.UpdateProductPhoto, OwnerAccess)
	productGroup.DELETE("/:productID", productController.DeleteProduct, OwnerAccess)

	server.Start(web)
}

package router

import (
	controllers "golang-template/controller"

	"github.com/labstack/echo/v4"
)

type Router struct {
	E *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	InitRouter(e)
	return &Router{
		E: e,
	}
}

func (r *Router) Start(addr string) error {
	r.E.Logger.Fatal(r.E.Start(addr))
	return nil
}

func InitRouter(e *echo.Echo) {
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	storeController := controllers.NewStoreController()

	authGroup := e.Group("/api/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
	authGroup.GET("/me", userController.GetUserByToken)
	authGroup.POST("/logout", userController.Logout)

	userGroup := e.Group("/api/user")
	// userGroup.Use(middleware.GetTokenNext)
	userGroup.PATCH("/", userController.UpdateUser)
	userGroup.DELETE("/", userController.DeleteUser)

	storeGroup := e.Group("/api/store")
	storeGroup.POST("/", storeController.CreateStore)
	storeGroup.GET("/", storeController.GetStoreByUserId)
	storeGroup.PATCH("/", storeController.UpdateStore)
	storeGroup.DELETE("/", storeController.DeleteStore)
}

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
	productController := controllers.NewProductController()
	healthGoalController := controllers.NewHealthGoalController()
	activityLevelController := controllers.NewActivityLevelController()
	productTypeController := controllers.NewProductTypeController()

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
	storeGroup.GET("/all", storeController.GetAllStores)

	productGroup := e.Group("/api/product")
	productGroup.POST("/", productController.CreateProduct)
	productGroup.GET("/:id", productController.GetProductById)
	productGroup.GET("/store/:id", productController.GetProductByStoreId)
	productGroup.GET("/", productController.GetAllProduct)
	productGroup.PATCH("/:id", productController.UpdateProduct)
	productGroup.DELETE("/:id", productController.DeleteProduct)
	productGroup.GET("/advertise/:id", productController.AdvertiseProduct)
	productGroup.GET("/unadvertise/:id", productController.UnadvertiseProduct)

	healthGoalGroup := e.Group("/api/health-goal")
	healthGoalGroup.POST("/", healthGoalController.CreateHealthGoal)
	healthGoalGroup.DELETE("/:id", healthGoalController.DeleteHealthGoal)
	healthGoalGroup.GET("/", healthGoalController.GetAllHealthGoal)

	activityLevelGroup := e.Group("/api/activity-level")
	activityLevelGroup.POST("/", activityLevelController.CreateActivityLevel)
	activityLevelGroup.DELETE("/:id", activityLevelController.DeleteActivityLevel)
	activityLevelGroup.GET("/", activityLevelController.GetAllActivityLevel)

	productTypeGroup := e.Group("/api/product-type")
	productTypeGroup.POST("/", productTypeController.CreateProductType)
	productTypeGroup.DELETE("/:id", productTypeController.DeleteProductType)
	productTypeGroup.GET("/", productTypeController.GetAllProductType)
}

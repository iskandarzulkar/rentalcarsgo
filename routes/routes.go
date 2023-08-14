package routes

import (
	"IskandarGo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// println(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("secret"),
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "REST API RENTAL CARS")
	})

	e.Static("/image", "./image")

	e.GET("/users", controllers.FectAllUsers)
	e.POST("/users", controllers.StoreUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	e.GET("generate-hash/:password", controllers.GenerateHashPassword)

	e.GET("/order", controllers.FectAllOrder)
	e.POST("/order", controllers.StoreOder)
	e.PUT("/order/:id", controllers.UpdateOrder)
	e.DELETE("/order/:id", controllers.DeleteOrder)

	e.GET("/cars", controllers.FectAllCars)
	e.POST("/cars", controllers.StoreCars)

	e.PUT("/cars/:id", controllers.UpdateCars)
	e.DELETE("/cars/:id", controllers.DeleteCars)

	e.POST("/image", controllers.ExStoreImage)

	e.GET("/carsorder", controllers.FetchCarsOrders)
	e.POST("login", controllers.CheckLogin)
	e.POST("loginUsers", controllers.CheckLoginUser)
	return e
}

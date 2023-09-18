package main

import (
	"github.com/dhikacruut/hris/controllers"
	"github.com/dhikacruut/hris/db"
	"github.com/dhikacruut/hris/middleware"
	"github.com/dhikacruut/hris/servers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
	db.CasbinAdapter()
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.POST("/signup", controllers.Signup)
		v1.POST("/login", controllers.Login)
		v1.GET("/validate", middleware.Authorize("resource", "*", db.Adapter), controllers.Validate)

		servers.InitializeRoutesEmployee(v1.Group("main"))
		servers.InitializeRoutesMpp(v1.Group("mpp"))
		servers.InitializeRoutesCompliance(v1.Group("compliance"))
		servers.InitializeRoutesService(v1.Group("service"))
		servers.InitializeRoutesPerformance(v1.Group("performance"))

		v1.POST("/logout", controllers.Logout)
	}

	router.Run()
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goserver/middlewares"
	"goserver/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//Middlewares
	router.Use(middlewares.Logger())

	// Routes
	router.GET("/", routes.Home)
	router.GET("/user", routes.User)

	// 404
	router.NoRoute(routes.NotFound)

	fmt.Println("Server started")
	router.Run(":8080")
}

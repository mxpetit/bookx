package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/handlers"
	"github.com/mxpetit/bookx/middleware"
)

func main() {
	app := gin.New()

	app.Use(middleware.Store())
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	bookGroup := app.Group("/book")
	{
		bookGroup.GET("", handlers.GetAllBooks)
		bookGroup.GET("/:id", handlers.GetBook)
		bookGroup.POST("", handlers.CreateBook)
	}

	app.Run(":8080")
}

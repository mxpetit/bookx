package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/handlers"
	"github.com/mxpetit/bookx/middleware"
)

// getRouter gets a new router with bookx's middlewares and routes.
func getRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Store())
	router.Use(middleware.Localisation())
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	bookGroup := router.Group("/book")
	{
		bookGroup.GET("", handlers.GetAllBooks)
		bookGroup.GET("/:id/previous", handlers.GetPreviousBooks)
		bookGroup.GET("/:id/next", handlers.GetNextBooks)
		bookGroup.GET("/:id", handlers.GetBook)
		bookGroup.POST("", handlers.CreateBook)
	}

	shelveGroup := router.Group("/shelve")
	{
		shelveGroup.GET("/:id", handlers.GetShelve)
		shelveGroup.POST("", handlers.CreateShelve)
	}

	return router
}

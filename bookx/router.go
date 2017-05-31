package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/handlers"
	"github.com/mxpetit/bookx/middleware"
	"net/http"
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

	authenticationGroup := router.Group("/authentication")
	{
		authenticationGroup.POST("", handlers.Authenticate)
		authenticationGroup.OPTIONS("", preflight)
	}

	return router
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://ui.book.xyz")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, X-Requested-With, Content-Type")
	c.JSON(http.StatusOK, struct{}{})
}

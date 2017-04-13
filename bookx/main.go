package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/handlers"
	"github.com/mxpetit/bookx/middleware"
	"github.com/nicksnyder/go-i18n/i18n"
	"os"
)

func main() {
	app := gin.New()

	i18n.MustLoadTranslationFile("translations/fr-FR.all.json")
	i18n.LoadTranslationFile("translations/en-US.all.json")

	app.Use(middleware.Store())
	app.Use(middleware.Localisation())
	app.Use(middleware.Cors())
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	bookGroup := app.Group("/book")
	{
		bookGroup.GET("", handlers.GetAllBooks)
		bookGroup.GET("/:id", handlers.GetBook)
		bookGroup.POST("", handlers.CreateBook)
	}

	app.Run(":" + getPort())
}

// getPort returns the value port contained in BOOKX_PORT or the default
// if any.
func getPort() string {
	port := os.Getenv("BOOKX_PORT")

	if port == "" {
		return "8080"
	}

	return port
}

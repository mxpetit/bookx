package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/handlers"
	"github.com/mxpetit/bookx/middleware"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	app := gin.New()

	i18n.MustLoadTranslationFile("translations/fr-FR.all.json")
	i18n.LoadTranslationFile("translations/en-US.all.json")

	app.Use(middleware.Store())
	app.Use(middleware.Localisation())
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

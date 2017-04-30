package main

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChannel
		stop()
		os.Exit(0)
	}()

	i18n.MustLoadTranslationFile("translations/fr-FR.all.json")
	i18n.LoadTranslationFile("translations/en-US.all.json")

	app := getRouter()
	app.Run(":" + getPort())
}

// getPort returns the value port contained in BOOKX_PORT or DEFAULT_PORT.
func getPort() string {
	port := os.Getenv("BOOKX_PORT")

	if port == "" {
		return DEFAULT_PORT
	}

	return port
}

// stop stops gracefully the running application.
func stop() {
	log.Println("Gracefully stopping...")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jaiiali/go-link-shortener/helpers"
	"github.com/jaiiali/go-link-shortener/internal/core/usecases"
	handlerLink "github.com/jaiiali/go-link-shortener/internal/handlers/link"
	repoLink "github.com/jaiiali/go-link-shortener/internal/repositories/link"
	"github.com/jaiiali/go-link-shortener/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	defer log.Sync() //nolint: errcheck

	// Repository
	linkRepo := repoLink.NewMemRepo(log)

	// UseCase
	todoUseCase := usecases.NewLinkUseCase(linkRepo, log)

	app := fiber.New()
	//app.Use(recover.New())
	api := app.Group("/api")

	// Handler
	handlerLink.NewHandler(todoUseCase, api)

	log.Info("Listening...")
	log.Panic(app.Listen(helpers.BuildAddr()))
}

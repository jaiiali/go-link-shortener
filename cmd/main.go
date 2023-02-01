package main

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jaiiali/go-link-shortener/helpers"
	"github.com/jaiiali/go-link-shortener/internal/core/usecases"
	handlerLink "github.com/jaiiali/go-link-shortener/internal/handlers/link"
	handlerRedirect "github.com/jaiiali/go-link-shortener/internal/handlers/redirect"
	repoLink "github.com/jaiiali/go-link-shortener/internal/repositories/link"
	"github.com/jaiiali/go-link-shortener/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	defer log.Sync() //nolint: errcheck

	// Repository
	linkRepo := repoLink.NewMemRepo()

	// UseCase
	linkUseCase := usecases.NewLinkUseCase(linkRepo, log)

	app := fiber.New()
	app.Use(fiberLogger.New())
	app.Use(fiberRecover.New())

	// Handler
	handlerRedirect.NewHandler(linkUseCase, app)

	apiGroup := app.Group("/api")
	handlerLink.NewHandler(linkUseCase, apiGroup)

	log.Info("Listening...")
	log.Panic(app.Listen(helpers.BuildAddr()))
}

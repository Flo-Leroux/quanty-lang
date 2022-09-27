package server

import (
	"github.com/Flo-Leroux/quanty-lang/internal/engine"
	"github.com/gofiber/fiber/v2"
)

func Prepare() *fiber.App {
	viewEngine := engine.
		New("./samples").
		Reload(true).
		Debug(true)

	app := fiber.New(fiber.Config{
		AppName:   "Quanty",
		GETOnly:   true,
		Immutable: true,
		Views:     viewEngine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main", nil)
	})

	app.Get("/faq", func(c *fiber.Ctx) error {
		return c.Render("faq", nil)
	})

	return app
}

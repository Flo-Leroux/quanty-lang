package server

import "github.com/gofiber/fiber/v2"

func handlePage(c *fiber.Ctx) error {
	_, err := c.WriteString("Hello")
	return err
}

func Prepare() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:   "Quanty",
		GETOnly:   true,
		Immutable: true,
	})

	app.Get("/", handlePage)

	return app
}

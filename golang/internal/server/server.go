package server

import (
	"github.com/Flo-Leroux/quanty-lang/internal/engine"
	"github.com/Flo-Leroux/quanty-lang/internal/language"
	"github.com/gofiber/fiber/v2"
)

func Run() *fiber.App {
	program, err := language.New(
		"Samples",
		language.WithDebug(),
		language.WithDirectory("./samples/pages"),
	)
	if err != nil {
		panic(err)
	}

	viewEngine := engine.
		New(program).
		Reload(true).
		Debug(true)

		// Default error handler
	var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {

		// Default 500 statuscode
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			code = e.Code

			if e.Code == fiber.StatusNotFound {
				return c.Redirect("/404")
			}
		}

		// Set Content-Type: text/plain; charset=utf-8
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		// Return statuscode with error message
		return c.Status(code).SendString(err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName:      "Quanty",
		GETOnly:      true,
		Immutable:    true,
		Views:        viewEngine,
		ErrorHandler: DefaultErrorHandler,
	})

	program.RegisterRoutes(app)

	return app
}

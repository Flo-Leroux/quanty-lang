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

	app := fiber.New(fiber.Config{
		AppName:   "Quanty",
		GETOnly:   true,
		Immutable: true,
		Views:     viewEngine,
	})

	allPaths := program.AllPaths()

	for _, name := range allPaths {
		route := "/" + name
		app.Get(route, func(c *fiber.Ctx) error {
			return c.Render(c.Route().Name, nil)
		}).Name(name)
	}

	return app
}

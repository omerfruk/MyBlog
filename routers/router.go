package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/handlers"
)

func Router(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.IndexRender)
	app.Get("/info", handlers.InfoRender)

	res := app.Group("/researcher/:key",handlers.ResearchersAll)                // /researcher
	res.Get("/computer", handlers.ResearchersAll)  // /researcher/Computer
	res.Get("/network", handlers.ResearchersAll)   // /researcher/Network
	res.Get("/user", handlers.ResearchersAll)      // /researcher/User
	res.Get("/developer", handlers.ResearchersAll) // /researcher/Developer
	res.Get("/hardware", handlers.ResearchersAll)  // /researcher/Hardware
	res.Get("/security", handlers.ResearchersAll)  // /researcher/Security

}

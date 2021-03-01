package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/handlers"
)

func Router(app *fiber.App) {

	app.Static("/", "./")
	// Initialize standard Go ht
	app.Get("/", handlers.IndexRender)
	app.Get("/info", handlers.InfoRender)

	res := app.Group("/researches")                // /researcher
	res.Get("/Computer", handlers.ResearchersAll)  // /researcher/Computer
	res.Get("/Network", handlers.ResearchersAll)   // /researcher/Network
	res.Get("/User", handlers.ResearchersAll)      // /researcher/User
	res.Get("/Developer", handlers.ResearchersAll) // /researcher/Developer
	res.Get("/Hardware", handlers.ResearchersAll)  // /researcher/Hardware
	res.Get("/Security", handlers.ResearchersAll)  // /researcher/Security

}

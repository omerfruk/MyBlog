package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/handlers"
)

func Router(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.IndexRender)
	app.Get("/info", handlers.InfoRender)
	app.Get("/login",handlers.Login)

	res := app.Group("/researcher/:key",handlers.ResearchersAll)                // /researcher
	res.Get("/Computer", handlers.ResearchersAll)  // /researcher/Computer
	res.Get("/User", handlers.ResearchersAll)      // /researcher/User
	res.Get("/Network", handlers.ResearchersAll)   // /researcher/Network
	res.Get("/Developer", handlers.ResearchersAll) // /researcher/Developer
	res.Get("/Hardware", handlers.ResearchersAll)  // /researcher/Hardware
	res.Get("/Security", handlers.ResearchersAll)  // /researcher/Security



}

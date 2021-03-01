package routers

import (
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/handlers"
)

func Router(app *fiber.App) {

	app.Get("/",handlers.ResearchersAll)
	app.Get("/info",handlers.ResearchersAll)
	res := app.Group("/researcher", handlers.ResearchersAll)// /researcher
	res.Get("/Computer", handlers.ResearchersAll)	// /researcher/Computer
	res.Get("/Network",handlers.ResearchersAll)	// /researcher/Network
	res.Get("/User",handlers.ResearchersAll)		// /researcher/User
	res.Get("/Developer",handlers.ResearchersAll)	// /researcher/Developer
	res.Get("/Hardware",handlers.ResearchersAll)	// /researcher/Hardware
	res.Get("/Security",handlers.ResearchersAll)	// /researcher/Security

	app.Static("/","./")
	// Initialize standard Go html template engine
	engine := html.New("./", ".html")

	app = fiber.New(fiber.Config{
		Views: engine,
	})


}


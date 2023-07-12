package router

import (
	"calculator/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", handlers.GetRoot)
	app.Get("/*", handlers.GetStatic)
	app.Post("/calculate", handlers.PostCalculate)
}

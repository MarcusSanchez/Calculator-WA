package main

import (
	"calculator/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := appInit()
	port := "3000"
	router.Router(app)
	app.Listen(":" + port)
}

func appInit() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		return c.Next()
	})
	return app
}

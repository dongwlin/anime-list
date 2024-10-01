package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App) {
	app.Use(recover.New())
	app.Use(cors.New())
}

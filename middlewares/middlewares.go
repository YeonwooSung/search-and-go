package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

/**
 * SetMiddlewares sets the middlewares for the application
 *
 * @param app *fiber.App
 */
func SetMiddlewares(app *fiber.App) {
	// Middleware for gzip compression
	app.Use(compress.New())
}

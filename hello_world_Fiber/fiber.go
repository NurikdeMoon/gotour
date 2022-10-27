package hello_world_Fiber

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello world!!!")

	})

	app.Listen(":2050")
}

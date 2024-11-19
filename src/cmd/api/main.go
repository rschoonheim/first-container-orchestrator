package main

import "github.com/gofiber/fiber/v2"

var (
	// runAddress - The address to run the API server on.
	runAddress = ":7331"

	// environment - The environment the application is running in.
	environment = "local"
)

func init() {

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Listen(runAddress)
}

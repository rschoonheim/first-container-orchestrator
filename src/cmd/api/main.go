package main

import (
	"github.com/gofiber/fiber/v3"
	"os"
)

var (
	// runAddress - The address to run the API server on.
	runAddress = ":7331"

	// environment - The environment the application is running in.
	environment = "local"
)

func init() {
	os.MkdirAll("storage/api", 0755)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(runAddress, fiber.ListenConfig{
		CertFile:       "certs/server.crt",
		CertKeyFile:    "certs/server.key",
		CertClientFile: "certs/server-ca-chain.crt",
	})
	if err != nil {
		panic(err)
	}

}

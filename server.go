package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", indexHandler)

	log.Fatalln(app.Listen(":3000"))
}

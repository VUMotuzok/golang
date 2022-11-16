package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func indexHandler(c *fiber.Ctx) error {
	connection := createConnection()
	log.Print(connection)
	return c.SendString("Hello")
}

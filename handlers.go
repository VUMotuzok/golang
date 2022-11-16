package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

var validate = validator.New()

type CreateFundsCreditHandlerDTO struct {
	UserId string `json:"userId" validate:"required,uuid"`
}

type CreateFundsReserveHandlerDTO struct {
	UserId string `json:"userId" validate:"required,uuid"`
}

type ApproveFundsReserveHandlerDTO struct {
	UserId string `json:"userId" validate:"required,uuid"`
}

func handleError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func createFundsCreditHandler(c *fiber.Ctx) error {
	connection := createConnection()
	log.Print(connection)

	dto := new(CreateFundsCreditHandlerDTO)

	if err := c.BodyParser(dto); err != nil {
		return handleError(c, err)
	}

	err := validate.Struct(dto)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendString(dto.UserId)
}

func createFundsReserveHandler(c *fiber.Ctx) error {
	connection := createConnection()
	log.Print(connection)

	dto := new(CreateFundsReserveHandlerDTO)

	if err := c.BodyParser(dto); err != nil {
		return handleError(c, err)
	}

	err := validate.Struct(dto)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendString(dto.UserId)
}

func approveFundsReserveHandler(c *fiber.Ctx) error {
	connection := createConnection()
	log.Print(connection)

	dto := new(ApproveFundsReserveHandlerDTO)

	if err := c.BodyParser(dto); err != nil {
		return handleError(c, err)
	}

	err := validate.Struct(dto)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendString(dto.UserId)
}

func getFundsHandler(c *fiber.Ctx) error {
	connection := createConnection()
	log.Print(connection)

	return c.SendString("Hello")
}

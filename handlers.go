package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"log"
	"service/models"
)

var validate = validator.New()

type CreateFundsCreditHandlerDTO struct {
	UserId string `json:"userId" validate:"required,uuid"`
	Amount int    `json:"amount" validate:"required,gte=0"`
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
	db := createConnection()

	dto := new(CreateFundsCreditHandlerDTO)

	if err := c.BodyParser(dto); err != nil {
		return handleError(c, err)
	}

	if err := validate.Struct(dto); err != nil {
		return handleError(c, err)
	}

	user := new(models.UserModel)

	exists, err := db.NewSelect().
		Model(user).
		Where("? = ?", bun.Ident("id"), dto.UserId).
		Exists(c.Context())

	if err != nil {
		return handleError(c, err)
	}

	if !exists {
		_, err = db.NewInsert().Model(&models.UserModel{
			Id:     uuid.Must(uuid.Parse(dto.UserId)),
			Amount: dto.Amount,
		}).Exec(c.Context())

		if err != nil {
			return handleError(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
	}

	err = db.NewSelect().
		Model(user).
		Where("? = ?", bun.Ident("id"), dto.UserId).
		Scan(c.Context())

	if err != nil {
		return handleError(c, err)
	}

	_, err = db.NewUpdate().
		Model(user).
		Where("? = ?", bun.Ident("id"), dto.UserId).
		Set("amount = ?", user.Amount+dto.Amount).
		Exec(c.Context())

	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
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

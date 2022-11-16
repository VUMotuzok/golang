package main

import (
	"database/sql"
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
	UserId    string `json:"userId" validate:"required,uuid"`
	ServiceId string `json:"serviceId" validate:"required,uuid"`
	OrderId   string `json:"orderId" validate:"required,uuid"`
	Amount    int    `json:"amount" validate:"required,gte=0"`
}

type ApproveFundsReserveHandlerDTO struct {
	UserId string `json:"userId" validate:"required,uuid"`
}

func handleError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func findUser(db *bun.DB, id string, c *fiber.Ctx) (*models.UserModel, bool, error) {
	user := new(models.UserModel)

	count, err := db.NewSelect().
		Model(user).
		Where("id=?", id).
		ScanAndCount(c.Context())

	if err != nil {
		return user, false, err
	}

	if count == 0 {
		return user, false, err
	}

	return user, true, err

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

	user, isFound, err := findUser(db, dto.UserId, c)
	if err != nil {
		return handleError(c, err)
	}

	if !isFound {
		_, err = db.NewInsert().Model(&models.UserModel{
			Id:     uuid.Must(uuid.Parse(dto.UserId)),
			Amount: dto.Amount,
		}).Exec(c.Context())

		if err != nil {
			return handleError(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
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
	db := createConnection()
	dto := new(CreateFundsReserveHandlerDTO)

	if err := c.BodyParser(dto); err != nil {
		return handleError(c, err)
	}

	if err := validate.Struct(dto); err != nil {
		return handleError(c, err)
	}

	user, isFound, err := findUser(db, dto.UserId, c)
	if err != nil {
		return handleError(c, err)
	}

	if !isFound || user.Amount < dto.Amount {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	tx, err := db.BeginTx(c.Context(), &sql.TxOptions{})

	_, err = tx.NewInsert().Model(&models.Transaction{
		Amount:    dto.Amount,
		OrderId:   uuid.Must(uuid.Parse(dto.OrderId)),
		ServiceId: uuid.Must(uuid.Parse(dto.ServiceId)),
		Status:    "reserved",
		UserId:    uuid.Must(uuid.Parse(dto.UserId)),
	}).Exec(c.Context())

	if err != nil {
		err = tx.Rollback()
		return handleError(c, err)
	}

	_, err = db.NewUpdate().
		Model(user).
		Where("? = ?", bun.Ident("id"), dto.UserId).
		Set("amount = ?", user.Amount-dto.Amount).
		Exec(c.Context())

	if err != nil {
		err = tx.Rollback()
		return handleError(c, err)
	}

	err = tx.Commit()

	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
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

	return c.SendStatus(fiber.StatusOK)
}

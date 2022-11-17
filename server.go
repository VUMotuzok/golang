package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
)

func createConnection() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/main?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	return db
}

func main() {
	app := fiber.New()

	app.Post("/funds/credit", createFundsCreditHandler)
	app.Post("/funds/reserve", createFundsReserveHandler)
	app.Post("/funds/reserve/approve", approveFundsReserveHandler)
	app.Get("/funds/:userId", getFundsHandler)

	log.Fatalln(app.Listen(":3000"))
}

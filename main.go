package main

import (
	controllerMain "ramaeqq/go-fiber-gorm/controllers"
	"ramaeqq/go-fiber-gorm/models/database"
	"ramaeqq/go-fiber-gorm/models/migration"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConfigDB()
	migration.DbMigration()

	app := fiber.New()

	controllerMain.Index(app)

	app.Listen(":9000")
}

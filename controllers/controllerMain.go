package controllerMain

import (
	"ramaeqq/go-fiber-gorm/config"
	"ramaeqq/go-fiber-gorm/handlers"
	"ramaeqq/go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

// func MiddlewareMain(ctm *fiber.Ctx) error {

// 	token := ctm.Get("x-token")
// 	if token != "secret" {
// 		return ctm.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"Message": "not Access",
// 		})
// 	}

// 	return ctm.Next()
// }

func Index(ctm *fiber.App) {

	ctm.Static("/public", config.ProjectRootPath+"/public/assets")

	ctm.Post("/login", handlers.LoginHandler)
	// ctm.Static("/public", "./public/assets")
	ctm.Get("/", middleware.Auth, handlers.MainHandlerGetAll)
	ctm.Get("/person/:id", handlers.MainHandlerGetById)
	ctm.Post("/", handlers.MainHandlerAdd)
	ctm.Put("/person/:id", handlers.MainHandlerUpdate)
	ctm.Delete("/person/:id", handlers.MainHandlerDelete)

}

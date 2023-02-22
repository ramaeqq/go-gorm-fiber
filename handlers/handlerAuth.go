package handlers

import (
	"log"
	"ramaeqq/go-fiber-gorm/models/database"
	"ramaeqq/go-fiber-gorm/models/entities"
	"ramaeqq/go-fiber-gorm/models/request"
	"ramaeqq/go-fiber-gorm/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(h *fiber.Ctx) error {

	auth := new(request.Auth)
	if err := h.BodyParser(auth); err != nil {
		return err
	}

	log.Println(auth)

	validate := validator.New()
	errValidate := validate.Struct(auth)
	if errValidate != nil {
		return h.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var person entities.Person
	err := database.DB.First(&person, "email = ?", auth.Email).Error
	if err != nil {
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "wrong credention",
		})
	}

	isValid := utils.CheckPasswordHash(auth.Password, person.Password)
	if !isValid {
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "wrong credential",
		})
	}

	//GENERATE JWT

	claims := jwt.MapClaims{}
	claims["name"] = person.Username
	claims["email"] = person.Email
	claims["address"] = person.Address
	claims["exp"] = time.Now().Add(time.Minute * 40).Unix()

	if person.Email == "rama@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGToken := utils.GenerateToken(&claims)
	if errGToken != nil {
		log.Println(errGToken)
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "wrong credential",
		})
	}

	return h.JSON(fiber.Map{
		"token": token,
	})

}

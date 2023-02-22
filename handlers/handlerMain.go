package handlers

import (
	"log"
	"ramaeqq/go-fiber-gorm/models/database"
	"ramaeqq/go-fiber-gorm/models/entities"
	"ramaeqq/go-fiber-gorm/models/request"
	"ramaeqq/go-fiber-gorm/models/response"
	"ramaeqq/go-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MainHandlerGetAll(h *fiber.Ctx) error {

	userInfo := h.Locals("userInfo")
	log.Println("user info data :: ", userInfo)

	var persons []entities.Person
	result := database.DB.Debug().Find(&persons)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return h.JSON(fiber.Map{
		"data": fiber.Map{
			"persons": &persons,
		},
	})

}

func MainHandlerAdd(h *fiber.Ctx) error {

	person := new(request.AddPerson)
	if err := h.BodyParser(person); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(person)
	if errValidate != nil {
		return h.Status(400).JSON(fiber.Map{
			"Message": "Failure",
			"error":   errValidate.Error(),
		})
	}

	addPerson := entities.Person{
		Username: person.Username,
		Email:    person.Email,
		Address:  person.Address,
	}

	hashedPassword, err := utils.HashingPassword(person.Password)
	if err != nil {
		log.Println(err)
		return h.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	addPerson.Password = hashedPassword

	errAdd := database.DB.Create(&addPerson).Error
	if errAdd != nil {
		return h.Status(500).JSON(fiber.Map{
			"Message": "Failure",
		})
	}

	return h.JSON(fiber.Map{
		"Message": "Success",
		"data":    addPerson,
	})

}

func MainHandlerGetById(h *fiber.Ctx) error {

	userId := h.Params("id")

	var person entities.Person
	err := database.DB.First(&person, "id = ?", userId).Error
	if err != nil {
		return h.Status(400).JSON(fiber.Map{
			"Message": "Failure | bad request",
		})
	}

	rMsg := response.MessageMain{
		Id:        person.Id,
		Username:  person.Username,
		Email:     person.Email,
		Address:   person.Address,
		Password:  person.Password,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	}

	return h.Status(200).JSON(fiber.Map{
		"Message": "Success",
		"Person":  rMsg,
	})

}

func MainHandlerUpdate(h *fiber.Ctx) error {

	personReq := new(request.UpdatePerson)
	if err := h.BodyParser(personReq); err != nil {
		return h.Status(400).JSON(fiber.Map{
			"Message": "Failure",
		})
	}

	var person entities.Person

	userId := h.Params("id")
	err := database.DB.First(&person, "id = ?", userId).Error
	if err != nil {
		return h.Status(404).JSON(fiber.Map{
			"message": "person not found",
		})
	}

	if personReq.Username != "" {
		person.Username = personReq.Username
	}
	person.Address = personReq.Address
	person.Password = personReq.Password
	errUpdate := database.DB.Save(&person).Error
	if errUpdate != nil {
		return h.Status(404).JSON(fiber.Map{
			"Message": "Not found",
		})
	}

	return h.JSON(fiber.Map{
		"Message": "Success",
	})

}

func MainHandlerDelete(h *fiber.Ctx) error {

	var person entities.Person

	userId := h.Params("id")
	err := database.DB.Debug().First(&person, "id = ?", userId).Error
	if err != nil {
		return h.Status(404).JSON(fiber.Map{
			"Message": "Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&person).Error

	if errDelete != nil {
		return h.JSON(fiber.Map{
			"Message": "failuer delete",
		})
	}

	return h.JSON(fiber.Map{
		"Message": "Success",
	})

}

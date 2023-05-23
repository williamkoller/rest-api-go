package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamkoller/rest-api-go/database"
	"github.com/williamkoller/rest-api-go/models"
)

func FindFact(c *fiber.Ctx) error {
	var facts []models.Fact

	result := database.DB.Db.Find(&facts)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(facts)
}

func FindFactById(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Fact

	result := database.DB.Db.Find(&fact, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fact)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

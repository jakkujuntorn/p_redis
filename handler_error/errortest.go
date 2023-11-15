package handlererror

import (
	"github.com/gofiber/fiber/v2"
)

func Error1(c *fiber.Ctx) error {
	response := fiber.Map{
		"Status":  NewResponseMessage("Error").Message,
		"Message": NewErrorMessage("Error Test Na ja").Message,
	}
	return c.JSON(response)
}

package russyhandler

import (
	"goredis/handler_error"
	"goredis/russy_services"
	
	"github.com/gofiber/fiber/v2"
)

type handlerData struct {
	ID       int
	Name     string
	Quantity int
}

type I_Handler interface {
	GetDataHandler(c *fiber.Ctx) error
}

type hanlerProduct struct {
	productService russyservices.I_Redis_Service
}

func NewHandler_Russy(productService russyservices.I_Redis_Service) I_Handler {
	return hanlerProduct{productService}
}

// GetDataHandler implements I_Handler
func (h hanlerProduct) GetDataHandler(c *fiber.Ctx) error {
	products, err := h.productService.GetRedisData()
	if err != nil {
		return handlererror.NewErrorMessage("Can not Connert Redis DB")
	}
	response := fiber.Map{
		"status":  handlererror.NewResponseMessage("Russy8"),
		"product": products,
	}

	return c.JSON(response)
}

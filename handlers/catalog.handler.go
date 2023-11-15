package handlers

import (
	"github.com/gofiber/fiber/v2"
	"goredis/services"
)

type catalogHandler struct {
	// ใช้ ของ services   ไม่ใช้ของ repository ****
	catalogSrv services.I_CatalogService
}

func NewCatalogHandler(catalogSrv services.I_CatalogService) ICatalogHandler {
	return catalogHandler{catalogSrv}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {

	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	// ใช้ fiber ปั้น response ใหม่ออกไป
	// เพื่อให้แยกระหว่าง service กับ handle 
	response := fiber.Map{
		"status":  "Code Bangkok",
		"product": products,
	}

	return c.JSON(response)
}

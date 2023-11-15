package handlers

import "github.com/gofiber/fiber/v2"

type ICatalogHandler interface {
	GetProducts(c *fiber.Ctx) error
}


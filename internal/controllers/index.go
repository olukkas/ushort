package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/olukkas/ushort/internal/controllers/templates"
)

func HelloRoute(c *fiber.Ctx) error {
	return templates.Index().Render(c.Context(), c.Response().BodyWriter())
}

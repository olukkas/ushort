package controllers

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/olukkas/ushort/internal/controllers/templates"
)

func HelloRoute(c *fiber.Ctx) error {
	c.Type("html")
	var buf bytes.Buffer
	if err := templates.Index().Render(c.Context(), &buf); err != nil {
		return err
	}
	return c.Send(buf.Bytes())
}

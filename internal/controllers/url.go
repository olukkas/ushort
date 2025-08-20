package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/olukkas/ushort/internal/controllers/templates"
	"github.com/olukkas/ushort/internal/repositories"
	"net/url"
)

type UrlController struct {
	urlRepo *repositories.UrlRepository
}

func NewUrlController(urlRepo *repositories.UrlRepository) *UrlController {
	return &UrlController{urlRepo: urlRepo}
}

func (u *UrlController) Shorten(ctx *fiber.Ctx) error {
	ctx.Type(fiber.MIMETextHTML)
	longUrl := ctx.FormValue("url")

	responseWriter := ctx.Response().BodyWriter()

	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return templates.ErrorMessage("Url inválida").
			Render(ctx.Context(), responseWriter)
	}

	entity := repositories.NewURL(longUrl)
	newUrl, err := u.urlRepo.Save(entity)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return templates.ErrorMessage(err.Error()).Render(ctx.Context(), responseWriter)
	}

	return templates.SuccessMessage(newUrl.Slug).Render(ctx.Context(), responseWriter)
}

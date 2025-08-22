package controllers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/olukkas/ushort/internal/repositories"
	"github.com/olukkas/ushort/internal/templates"
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
		return templates.ErrorMessage("Url inválida").
			Render(ctx.Context(), responseWriter)
	}

	entity := repositories.NewURL(longUrl)
	newUrl, err := u.urlRepo.Save(entity)
	if err != nil {
		return templates.ErrorMessage(err.Error()).Render(ctx.Context(), responseWriter)
	}

	return templates.SuccessMessage(ctx.BaseURL()+"/"+newUrl.Slug).Render(ctx.Context(), responseWriter)
}

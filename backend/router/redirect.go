package router

import (
	"link-shortner/service"
	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) {
	shortUrl := c.Params("shortUrl")
	url := service.GetUrl(shortUrl)
	c.JSON(fiber.Map{"url": url})
}
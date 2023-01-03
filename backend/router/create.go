package router

import (
	"link-shortner/service"
	"github.com/gofiber/fiber/v2"
) 

func CreateLink(c *fiber.Ctx) {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(503).SendString(err.Error())
		return
	}

	shortenUrl := service.CreateShortenUrl(data["url"])
	
	c.JSON(fiber.Map{
		"shortUrl": shortenUrl,
	})
}

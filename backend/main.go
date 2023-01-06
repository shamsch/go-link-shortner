package main

import (
	"fmt"
	"link-shortner/db"
	"link-shortner/router"

	"github.com/gofiber/fiber/v2"
)	

func main() {
	app := fiber.New()
	db.InitDB()
	app.Post("/create", func(c *fiber.Ctx) error {
		fmt.Println("Creating a link")
		router.CreateLink(c)
		return nil
	})
	app.Get("/:shortUrl", func(c *fiber.Ctx) error {
		fmt.Println("Getting the long link")
		router.Redirect(c)
		return nil
	})
	fmt.Println("Listening on port 3000")
	app.Listen(":3000")
}

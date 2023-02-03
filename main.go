package main

import (
	"log"

	"github.com/HoangYen2002/bookstore/migrates"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my webs222ite!")
}

func main() {
	migrates.ConnectDb()
	app := fiber.New()
	app.Get("/", welcome)
	log.Fatalln(app.Listen(":8081"))
}

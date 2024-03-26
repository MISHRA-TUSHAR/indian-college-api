package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()
	handlerObj := handlers.NewAPIHandler()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	app.Get("colleges/", handlerObj.SearchCollege)
	app.Get("colleges/states", handlerObj.GetAllStates)
	app.Get("colleges/:state/districts", handlerObj.GetDistrictsByState)
	app.Get("colleges/:state", handlerObj.GetAllCollegesInState)
	app.Get("colleges/:state/:district", handlerObj.GetAllCollegesInDistrict)

	log.Fatal(app.Listen(":3000"))
}

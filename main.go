package main

import (
	"episode-3/app/controller"
	"episode-3/app/model"
	"episode-3/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	services.InitDatabase()
	db := services.DB
	db.AutoMigrate(&model.Movie{})

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("hello world")
	})

	api := app.Group("/api/v1/movies")

	api.Get("/", controller.GetMovie)
	api.Get("/:id", controller.GetMovieID)
	api.Post("/", controller.PostMovie)
	api.Put("/:id", controller.PutMovie)
	api.Delete("/:id", controller.DeleteMovie)

	log.Fatal(app.Listen(":8081"))
}

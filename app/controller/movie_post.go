package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"

	"github.com/gofiber/fiber/v2"
)

func PostMovie(c *fiber.Ctx) error {
	db := services.DB
	var movieApi model.MovieAPI
	if err := c.BodyParser(&movieApi); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	movie := &model.Movie{MovieAPI: movieApi}
	db.Create(&movie)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    movie,
	})

}

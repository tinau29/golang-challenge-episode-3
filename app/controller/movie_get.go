package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetMovie(c *fiber.Ctx) error {
	db := services.DB
	var movie []model.Movie
	db.Model(&model.Movie{}).Find(&movie)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    movie,
	})
}

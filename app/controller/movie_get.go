package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetMovie show all movies in array
// @Summary show all movies in array
// @Description show all movies in array
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Movie{} "OK"
// @Router /movies [get]
// @Tags Movie
func GetMovie(c *fiber.Ctx) error {
	db := services.DB
	var movie []model.Movie
	result := db.Model(&model.Movie{}).Find(&movie)
	log.Println(result)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    movie,
	})
}

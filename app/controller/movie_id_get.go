package controller

import (
	"episode-3/app/libraries"
	"episode-3/app/model"
	"episode-3/app/services"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetMovieID godoc
// @Summary Get a movie by ID
// @Description Get a movie by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Movie{} "OK"
// @Success 400
// @Success 404
// @Param id path string true "Movie ID"
// @Router /movies/{id} [get]
// @Tags Movie
func GetMovieID(c *fiber.Ctx) error {
	db := services.DB

	id := c.Params("id")
	if !libraries.RegexCheckNumeric(id) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid id format",
		})
	}

	var movie model.Movie
	if err := db.Where(`id = ?`, id).First(&movie).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    movie,
	})
}

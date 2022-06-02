package controller

import (
	"episode-3/app/libraries"
	"episode-3/app/model"
	"episode-3/app/services"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DeleteMovie godoc
// @Summary Delete movie by ID
// @Description Delete movie by ID
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Movie ID"
// @Router /movies/{id} [delete]
// @Tags Movie
func DeleteMovie(c *fiber.Ctx) error {
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

	db.Where(`id = ?`, id).Delete(&movie)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

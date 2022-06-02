package controller

import (
	"episode-3/app/libraries"
	"episode-3/app/model"
	"episode-3/app/services"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PostMovie godoc
// @Summary Update movie by ID
// @Description Update movie ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Movie{} "OK"
// @Success 400
// @Success 404
// @Param id path string true "Movie ID"
// @Param data body model.MovieAPI true "Movie data"
// @Router /movies/{id} [put]
// @Tags Movie
func PutMovie(c *fiber.Ctx) error {
	db := services.DB
	var movieApi model.MovieAPI

	id := c.Params("id")
	if !libraries.RegexCheckNumeric(id) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid id format",
		})
	}

	if err := c.BodyParser(&movieApi); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var movieExist model.Movie
	if err := db.Where(`id = ?`, id).First(&movieExist).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	movie := &model.Movie{MovieAPI: movieApi}

	title := movieExist.Title
	if movie.Title != "" {
		title = movie.Title
	}

	year := movieExist.Year
	if movie.Year != 0 {
		year = movie.Year
	}

	summary := movieExist.Summary
	if movie.Summary != "" {
		summary = movie.Summary
	}

	director := movieExist.Director
	if movie.Director != "" {
		director = movie.Director
	}

	movie.ID = movieExist.ID
	movie.SetMovie(title, year, summary, director)
	db.Where(`id = ?`, id).Updates(&movie)

	message := fmt.Sprintf(`movie with  id %s has been updated`, id)

	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    movie,
	})
}

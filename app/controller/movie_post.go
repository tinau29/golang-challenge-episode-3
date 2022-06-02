package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostMovie godoc
// @Summary Create new movie
// @Description Create new movie
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Movie{} "OK"
// @Success 400
// @Param data body model.MovieAPI true "Movie data"
// @Router /movies [post]
// @Tags Movie
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

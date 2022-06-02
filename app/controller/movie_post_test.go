package controller

import (
	"bytes"
	"episode-3/app/model"
	"episode-3/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostMovie(t *testing.T) {

	app := fiber.New()
	app.Post("/api/v1/movies", PostMovie)

	payload := bytes.NewReader([]byte(`
	{ 
		"title": "Movie AB",
		"year": 2021,
		"summary": "summary",
		"director": "test ok"
	}
	`))

	var movieApi model.MovieAPI
	movieApi.Title = "Test Title"
	movieApi.Summary = "Test Summary"
	movieApi.Year = 2020
	movieApi.Director = "Joni"

	movie := &model.Movie{MovieAPI: movieApi}

	db := services.InitDatabaseTest()
	db.Create(&movie)

	req, _ := http.NewRequest("POST", "/api/v1/movies", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid request")

	req, _ = http.NewRequest("POST", "/api/v1/movies", payload)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

}

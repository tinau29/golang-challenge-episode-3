package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetMovie(t *testing.T) {

	type TestGetMovieStruct struct {
		description  string
		route        string
		expectedCode int
	}

	caseTestGetMovie := []TestGetMovieStruct{
		{
			description:  "get response",
			route:        "/api/v1/movies",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/api/v1/movies", GetMovie)

	var movie model.Movie
	movie.Title = "Test Title"
	movie.Summary = "Test Summary"
	movie.Year = 2020
	movie.Director = "Joni"

	db := services.InitDatabaseTest()
	db.Create(&movie)

	for _, test := range caseTestGetMovie {
		req, _ := http.NewRequest("GET", test.route, nil)
		req.Header.Set("accept", "application/json")
		res, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "send request")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)
	}
}

package controller

import (
	"episode-3/app/model"
	"episode-3/app/services"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteMovie(t *testing.T) {

	type TestDeleteMovieStruct struct {
		description  string
		route        string
		expectedCode int
	}

	app := fiber.New()

	app.Delete("/api/v1/movies/:id", DeleteMovie)

	var movie model.Movie
	movie.Title = "Test Title"
	movie.Summary = "Test Summary"
	movie.Year = 2020
	movie.Director = "Joni"

	db := services.InitDatabaseTest()
	db.Create(&movie)
	id := strconv.Itoa(*&movie.ID)

	caseTest := []TestDeleteMovieStruct{
		{
			description:  "get response",
			route:        "/api/v1/movies/" + id,
			expectedCode: 200,
		},
		{
			description:  "get response",
			route:        "/api/v1/movies/test",
			expectedCode: 400,
		},
		{
			description:  "get response",
			route:        "/api/v1/movies/0",
			expectedCode: 404,
		},
	}

	for _, test := range caseTest {
		req, _ := http.NewRequest("DELETE", test.route, nil)
		req.Header.Set("accept", "application/json")
		res, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "send request")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)
	}
}

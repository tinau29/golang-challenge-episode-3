package controller

import (
	"bytes"
	"episode-3/app/model"
	"episode-3/app/services"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutMovie(t *testing.T) {

	app := fiber.New()
	app.Put("/api/v1/movies/:id", PutMovie)

	var movie model.Movie
	movie.Title = "Test Title"
	movie.Summary = "Test Summary"
	movie.Year = 2020
	movie.Director = "Joni"

	var movieDeleted model.Movie
	movieDeleted.Title = "Test Title"
	movieDeleted.Summary = "Test Summary"
	movieDeleted.Year = 2020
	movieDeleted.Director = "Joni"

	db := services.InitDatabaseTest()
	db.Create(&movie)
	id := strconv.Itoa(*&movie.ID)

	db.Create(&movieDeleted)
	idDeleted := strconv.Itoa(*&movieDeleted.ID)
	db.Where(`id = ?`, idDeleted).Delete(&movieDeleted)

	var movieExist model.Movie
	db.Where(`id = ?`, id).First(&movieExist)

	payload := bytes.NewReader([]byte(`
	{ 
		"title": "Movie AB",
		"year": 2021,
		"summary": "summary",
		"director": "test ok"
	}
	`))

	req, _ := http.NewRequest("PUT", "/api/v1/movies/"+id, payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	req, _ = http.NewRequest("PUT", "/api/v1/movies/"+id, nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad body parser")

	req, _ = http.NewRequest("PUT", "/api/v1/movies/test", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid id format")

	req, _ = http.NewRequest("PUT", "/api/v1/movies/"+idDeleted, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "not found")

}

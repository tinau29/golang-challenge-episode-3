package controller

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteMovie(t *testing.T) {
	app := fiber.New()

	app.Delete("/api/v1/movies", DeleteMovie)

	req, _ := http.NewRequest("DELETE", "/api/v1/movies/3", nil)
	// req.Header.Set("accept", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get response")

	req, _ = http.NewRequest("DELETE", "/api/v1/movies/2", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "get response")

	req, _ = http.NewRequest("DELETE", "/api/v1/movies/p", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "get response")
}

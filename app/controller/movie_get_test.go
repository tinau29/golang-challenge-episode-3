package controller

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetMovie(t *testing.T) {
	app := fiber.New()

	app.Get("/api/v1/movies", GetMovie)

	req, _ := http.NewRequest("GET", "/api/v1/movies", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get response")

	req, _ = http.NewRequest("GET", "/api/v1/movies", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "get response")
}

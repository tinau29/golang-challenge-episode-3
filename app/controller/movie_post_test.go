package controller

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostMovie(t *testing.T) {
	app := fiber.New()

	payload := bytes.NewReader([]byte(`
		{
			"title": "Movie AB",
			"year": 2020,
			"summary": "cek aja",
			"director: "test ok"
		}
	`))

	req, _ := http.NewRequest("POST", "/api/v1/movies", payload)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")
}

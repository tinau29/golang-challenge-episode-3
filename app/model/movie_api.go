package model

// MovieAPI model
type MovieAPI struct {
	Title    string `json:"title,omitempty" example:"Movie Title"`
	Year     int    `json:"year,omitempty" example:"2021"`
	Summary  string `json:"summary,omitempty" example:"ok"`
	Director string `json:"director,omitempty" example:"joni"`
}

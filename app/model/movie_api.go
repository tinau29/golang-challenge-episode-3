package model

type MovieAPI struct {
	Title    string `json:"title, omitempty"`
	Year     int    `json:"year, omitempty"`
	Summary  string `json:"summary, omitempty"`
	Director string `json:"director, omitempty"`
}

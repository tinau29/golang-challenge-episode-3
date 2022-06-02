package model

// Movie model
type Movie struct {
	ID int `json:"id" example:"1"`
	MovieAPI
}

func (m *Movie) SetMovie(title string, year int, summary string, director string) {
	m.Title = title
	m.Year = year
	m.Summary = summary
	m.Director = director
}

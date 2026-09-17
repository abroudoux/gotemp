package film

import (
	_ "embed"
	"encoding/json"

	"github.com/abroudoux/gotemp/internal/domain"
)

//go:embed films.json
var filmsData []byte

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindAll() ([]domain.Film, error) {
	var films []domain.Film
	if err := json.Unmarshal(filmsData, &films); err != nil {
		return nil, err
	}
	return films, nil
}

package domain

type Film struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

type FilmRepository interface {
	FindAll() ([]Film, error)
}

type FilmService struct {
	repo FilmRepository
}

func NewFilmService(repo FilmRepository) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) ListFilms() ([]Film, error) {
	return s.repo.FindAll()
}

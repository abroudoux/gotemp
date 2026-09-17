package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/abroudoux/gotemp/internal/domain"
)

type FilmHandler struct {
	service *domain.FilmService
}

func NewFilmHandler(service *domain.FilmService) *FilmHandler {
	return &FilmHandler{service: service}
}

func (h *FilmHandler) ListFilms(c *gin.Context) {
	films, err := h.service.ListFilms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, films)
}

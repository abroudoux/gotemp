package http

import "github.com/gin-gonic/gin"

func NewRouter(filmHandler *FilmHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/films", filmHandler.ListFilms)

	return router
}

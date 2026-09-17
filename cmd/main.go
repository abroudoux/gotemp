package main

import (
	"log"

	adapterhttp "github.com/abroudoux/gotemp/internal/adapters/http"
	"github.com/abroudoux/gotemp/internal/domain"
	"github.com/abroudoux/gotemp/internal/infrastructure/config"
	"github.com/abroudoux/gotemp/internal/infrastructure/film"
)

func main() {
	cfg := config.Load()

	filmRepo := film.NewRepository()
	filmService := domain.NewFilmService(filmRepo)
	filmHandler := adapterhttp.NewFilmHandler(filmService)

	router := adapterhttp.NewRouter(filmHandler)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

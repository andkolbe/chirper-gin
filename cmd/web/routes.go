package main

import (
	"github.com/andkolbe/chirper-gin/internal/handlers"
)

func initializeRoutes() {
	router.GET("/", handlers.HomePage)

	chirpRoutes := router.Group("/chirps")
	{
		chirpRoutes.GET("/:chirp_id", handlers.GetChirp)
	}
}


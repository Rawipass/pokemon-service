package routes

import (
	"github.com/Rawipass/pokemon-service/internal/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon/:name", http.GetPokemonHandler)
	router.GET("/pokemon/:name/ability", http.GetAbilitiesHandler)
	router.GET("/pokemon/random", http.GetRandomPokemonHandler)

	return router
}

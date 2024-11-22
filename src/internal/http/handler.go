package http

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Rawipass/pokemon-service/pkg/pokeapi"
	"github.com/gin-gonic/gin"
)

func GetPokemonHandler(c *gin.Context) {
	name := c.Param("name")
	data, err := pokeapi.FetchPokemon(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetAbilitiesHandler(c *gin.Context) {
	name := c.Param("name")
	data, err := pokeapi.FetchPokemon(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"abilities": data.Abilities})
}

func GetRandomPokemonHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	randomID := 1 + rand.Intn(1010)
	data, err := pokeapi.FetchPokemonByID(randomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch random Pokemon"})
		return
	}
	c.JSON(http.StatusOK, data)
}

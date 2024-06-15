package http

import (
	"net/http"

	"github.com/api-pokemon/internal/adapters/domain"
	usecase "github.com/api-pokemon/internal/adapters/useCase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.Pokemon
}

func NewPokemonHandler(us usecase.Pokemon) *Handler {
	return &Handler{usecase: us}
}

func (h *Handler) Getpokemon(c *gin.Context) {
	var pokemonName domain.Pokemon
	name := c.Param("name")
	pokemonName.Name = name
	pok, err := h.usecase.Get(pokemonName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pok,
	})
}

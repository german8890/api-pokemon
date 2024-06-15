package repository

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/api-pokemon/internal/adapters/domain"
)

type Repository interface {
	GetPokemonByName(name string) (*domain.PokemonResponse, error)
}

type repository struct {
	client *http.Client
}

func NewRepository(client *http.Client) Repository {
	return &repository{client: client}
}

func (r *repository) GetPokemonByName(name string) (*domain.PokemonResponse, error) {
	resp, err := r.client.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pokemon *domain.PokemonResponse

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	return pokemon, nil

}

package usecase

import (
	"github.com/api-pokemon/internal/adapters/domain"
	"github.com/api-pokemon/internal/adapters/repository"
)

type Pokemon interface {
	Get(domain.Pokemon) (*domain.PokemonResponse, error)
}

type usecase struct {
	repo repository.Repository
}

func NewUse(repo repository.Repository) *usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Get(pokemon domain.Pokemon) (*domain.PokemonResponse, error) {
	resp, err := u.repo.GetPokemonByName(pokemon.Name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

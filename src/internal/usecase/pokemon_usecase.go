package usecase

import (
	"github.com/Rawipass/pokemon-service/internal/repository"
)

type PokemonUsecase struct {
	repo repository.PokemonRepository
}

func NewPokemonUseCase() *PokemonUsecase {
	usecase := PokemonUsecase{}
	return &usecase
}

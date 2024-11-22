package repository

import "database/sql"

type PokemonRepository struct {
	DB *sql.DB
}

func NewPokemonRepository() *PokemonRepository {
	repo := PokemonRepository{}
	return &repo
}

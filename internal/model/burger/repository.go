package burger

import (
	"burger-api/internal/model"
)

// normally ctx should also be provided, but its omitted to keep the api cleaner
type Repository interface {
	FindAll(page model.Page, nameEq string) ([]Burger, error)
	FindOne(id string) (*Burger, error)
	FindRandom() (*Burger, error)
	Save(burger *Burger) (*Burger, error)
}

var Repo Repository

func InitRepository(repository Repository) {
	Repo = repository
}
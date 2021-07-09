package repository

//go:generate mockgen -destination=./mocks/mock_repository.go -package=mocks github.com/ceperapl/app-poc/pkg/repository RecipeRepository

import (
	"github.com/ceperapl/app-poc/pkg/models"
)

// RecipeRepository is an abstract repository layer for data access
type RecipeRepository interface {
	CreateRecipe(recipe *models.Recipe) error
	GetRecipe(id string) (*models.Recipe, error)
	ListRecipes(filterBy string, sortBy string, limit int, page int) (recipes []models.Recipe, count int, err error)
	UpdateRecipe(recipe *models.Recipe) error
	DeleteRecipe(id string) error
}

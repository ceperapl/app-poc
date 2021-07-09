package usecase

//go:generate mockgen -destination=./mocks/mock_service.go -package=mocks github.com/ceperapl/app-poc/pkg/usecase RecipeService

import (
	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/repository"
)

// RecipeService is a service for managing recipes
type RecipeService interface {
	CreateRecipe(recipe *models.Recipe) (*models.Recipe, error)
	GetRecipe(id string) (*models.Recipe, error)
	ListRecipes(filterBy string, sortBy string,
		limit int, page int) (recipes []models.Recipe, count int, err error)
	UpdateRecipe(recipe *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(cid string) error
}

// NewRecipeService returns a recipe service
func NewRecipeService(repos repository.RecipeRepository) RecipeService {
	return &recipeService{repos}
}

type recipeService struct {
	repos repository.RecipeRepository
}

func (ts *recipeService) CreateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	if err := ts.repos.CreateRecipe(recipe); err != nil {
		return nil, err
	}
	return recipe, nil
}

func (ts *recipeService) GetRecipe(id string) (*models.Recipe, error) {
	return ts.repos.GetRecipe(id)
}

func (ts *recipeService) ListRecipes(filterBy string, sortBy string,
	limit int, page int) (recipes []models.Recipe, count int, err error) {
	return ts.repos.ListRecipes(filterBy, sortBy, limit, page)
}

func (ts *recipeService) UpdateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	if err := ts.repos.UpdateRecipe(recipe); err != nil {
		return nil, err
	}
	return recipe, nil
}

func (ts *recipeService) DeleteRecipe(id string) error {
	return ts.repos.DeleteRecipe(id)
}

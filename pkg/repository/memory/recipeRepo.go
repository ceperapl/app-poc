package memory

import (
	"fmt"
	"time"

	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/repository"
	"github.com/google/uuid"
)

// RecipeRepo implements recipe repo in memory
type RecipeRepo struct {
	recipes []models.Recipe
}

// NewRecipeRepo returns memory repository implementation
func NewRecipeRepo() (*RecipeRepo, error) {
	return &RecipeRepo{recipes: []models.Recipe{}}, nil
}

// CreateRecipe creates new recipe in memory if ID is not exists
func (m *RecipeRepo) CreateRecipe(recipe *models.Recipe) error {
	if recipe.ID == "" {
		recipe.ID = uuid.New().String()
	} else if _, err := uuid.Parse(recipe.ID); err != nil {
		return repository.ErrInvalidUUID
	}
	now := time.Now()
	recipe.CreatedAt = now
	recipe.UpdatedAt = now
	for _, t := range m.recipes {
		if t.ID == recipe.ID {
			return fmt.Errorf("Recipe ID already exists: %s", recipe.ID)
		}
	}
	m.recipes = append(m.recipes, *recipe)
	return nil
}

// GetRecipe returns recipe by id
func (m *RecipeRepo) GetRecipe(id string) (*models.Recipe, error) {
	for _, t := range m.recipes {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, repository.ErrRecipeNotFound
}

// ListRecipes returns recipes
func (m *RecipeRepo) ListRecipes(filterBy string, sortBy string, limit int,
	page int) (recipes []models.Recipe, count int, err error) {
	return m.recipes, len(m.recipes), nil
}

// UpdateRecipe updates recipe
func (m *RecipeRepo) UpdateRecipe(recipe *models.Recipe) error {
	for i, t := range m.recipes {
		if t.ID == recipe.ID {
			m.recipes[i] = *recipe
			return nil
		}
	}
	return repository.ErrRecipeNotFound
}

// DeleteRecipe deletes recipe
func (m *RecipeRepo) DeleteRecipe(id string) error {
	for i, t := range m.recipes {
		if t.ID == id {
			m.recipes = append(m.recipes[:i], m.recipes[i+1:]...)
			return nil
		}
	}
	return nil
}

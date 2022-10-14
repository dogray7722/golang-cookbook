package recipes

import (
	"github.com/golang-cookbook/utils/errors"
)

// GetRecipe returns an individual recipe by recipe id
func (recipe *Recipe) GetRecipe() *errors.RestErr {
	return nil
}

// SaveRecipe adds a new recipe
func (recipe *Recipe) SaveRecipe() *errors.RestErr {
	return nil
}

// DeleteRecipe removes a recipe 
func (recipe *Recipe) DeleteRecipe(recipeId int64) *errors.RestErr {
	return nil
}

// ListRecipes returns a set of recipes
func (recipe *Recipe) ListRecipes() ([]Recipe, *errors.RestErr) {
	var recipes []Recipe
	recipes = append(recipes, *recipe)

		
	return recipes, nil
}

// UpdateRecipe updates an individual recipe
func (recipe *Recipe) UpdateRecipe() *errors.RestErr {
	return nil
}

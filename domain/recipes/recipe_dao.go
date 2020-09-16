package recipes

import (
	"fmt"
	"github.com/dogray7722/golang-cookbook/utils/errors"
)

var (
	recipesDB = make(map[int64]*Recipe)
)

func (recipe *Recipe) Get() *errors.RestErr {
	result := recipesDB[recipe.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("recipe %d not found", recipe.Id))
	}

	recipe.Id = result.Id
	recipe.Ingredients = result.Ingredients
	recipe.Instructions = result.Instructions
	recipe.DateCreated = result.DateCreated

	return nil
}

func (recipe *Recipe) Save() *errors.RestErr {
	current := recipesDB[recipe.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("recipe %d already exists", recipe.Id))
	}

	if current.Name == recipe.Name {
		return errors.NewBadRequestError(fmt.Sprintf("recipe name %d already exists", recipe.Name))
	}

	recipesDB[recipe.Id] = recipe
	return nil
}

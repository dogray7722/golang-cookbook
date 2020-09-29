package recipes

import (
	"fmt"
	"github.com/dogray7722/golang-cookbook_dogray7722/datasources/postgres/recipes_db"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/date_utils"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
)

var (
	recipesDB = make(map[int64]*Recipe)
)

func (recipe *Recipe) Get() *errors.RestErr {
	newClient := recipes_db.Client
	if err := newClient.Ping(); err != nil {
		panic(err)
	}

	result := recipesDB[recipe.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("recipe %d not found", recipe.Id))
	}

	recipe.Id = result.Id
	recipe.Name = result.Name
	recipe.Ingredients = result.Ingredients
	recipe.Instructions = result.Instructions
	recipe.DateCreated = result.DateCreated
	recipe.Status = result.Status

	return nil
}

func (recipe *Recipe) Save() *errors.RestErr {
	current := recipesDB[recipe.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("recipe %d already exists", recipe.Id))
	}

	recipe.DateCreated = date_utils.GetNowString()

	recipesDB[recipe.Id] = recipe
	return nil
}

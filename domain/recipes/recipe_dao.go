package recipes

import (
	"fmt"
	"github.com/dogray7722/golang-cookbook_dogray7722/datasources/postgres/recipes_db"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
)

const (
	queryInsertRecipe     = "INSERT INTO recipes(name, instructions, date_created, status) VALUES($1, $2, $3, $4);"
	queryInsertIngredient = "INSERT INTO ingredients(serving_size, item, date_created) VALUES($1, $2, $3);"
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
	stmt, err := recipes_db.Client.Prepare(queryInsertRecipe)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(recipe.Name, recipe.Ingredients, recipe.Instructions, recipe.DateCreated, recipe.Status)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}
	recipeId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}

	recipe.Id = recipeId

	return nil
}

func (recipe *Recipe) SaveIngredients() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryInsertIngredient)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	for i := range recipe.Ingredients {
		insertResult, err := stmt.Exec(recipe.Ingredients[i].ServingSize, recipe.Ingredients[i].Item)
		if err != nil {
			return errors.NewInternalServerError(
				fmt.Sprintf("failed to save ingredient: %s", err.Error()))
		}
		ingredientId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.NewInternalServerError(
				fmt.Sprintf("failed to save ingredient: %s", err.Error()))
		}

		recipe.Ingredients[i].Id = ingredientId
		//also need to insert the values into the lookup table
		//can call that from here

	}

	return nil
}

package recipes

import (
	"fmt"
	"github.com/dogray7722/golang-cookbook_dogray7722/datasources/postgres/recipes_db"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
	"strings"
)

const (
	queryInsertRecipe     = "INSERT INTO recipes(name, instructions, status) VALUES($1, $2, $3) RETURNING id;"
	queryInsertIngredient = "INSERT INTO ingredients(serving_size, item) VALUES($1, $2) RETURNING id;"
	queryInsertLookup     = "INSERT INTO recipes_to_ingredients(recipe_id, ingredient_id) VALUES($1, $2);"
	indexUniqueRecipeName = "constraint_name"
)

//var (
//	recipesDB = make(map[int64]*Recipe)
//)

//func (recipe *Recipe) Get() *errors.RestErr {
//	newClient := recipes_db.Client
//	if err := newClient.Ping(); err != nil {
//		panic(err)
//	}
//
//	result := recipesDB[recipe.Id]
//	if result == nil {
//		return errors.NewNotFoundError(fmt.Sprintf("recipe %d not found", recipe.Id))
//	}
//
//	recipe.Id = result.Id
//	recipe.Name = result.Name
//	recipe.Ingredients = result.Ingredients
//	recipe.Instructions = result.Instructions
//	recipe.DateCreated = result.DateCreated
//	recipe.Status = result.Status
//
//	return nil
//}

func (recipe *Recipe) Save() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryInsertRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(recipe.Name, recipe.Instructions, recipe.Status).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueRecipeName) {
			return errors.NewBadRequestError(fmt.Sprintf(
				"recipe name %s already exists", recipe.Name))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}

	recipe.Id = id

	return nil
}

func (recipe *Recipe) SaveIngredients() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryInsertIngredient)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare ingredients query: %s", err.Error()))
	}
	defer stmt.Close()

	for i := range recipe.Ingredients {
		var id int64
		err := stmt.QueryRow(recipe.Ingredients[i].ServingSize, recipe.Ingredients[i].Item).Scan(&id)
		if err != nil {
			return errors.NewInternalServerError(
				fmt.Sprintf("failed to save ingredient: %s", err.Error()))
		}

		recipe.Ingredients[i].Id = id
		if err := saveRecipeIngredient(recipe.Id, id); err != nil {
			return errors.NewInternalServerError(
				fmt.Sprintf("failed to save ingredient: %s", err))
		}
	}

	return nil
}

func saveRecipeIngredient(recipeId, ingredientId int64) *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryInsertLookup)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare recipe to ingredients query: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipeId, ingredientId)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save recipes to ingredients: %s", err.Error()))
	}

	return nil
}

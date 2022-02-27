package recipes

import (
	"fmt"
	"strings"

	"github.com/golang-cookbook/datasources/postgres/recipes_db"
	"github.com/golang-cookbook/utils/errors"
	"github.com/lib/pq"
)

const (
	indexUniqueRecipeTitle = "constraint_title"
	errorNoRows           = "no rows in result set"

	queryInsertRecipe = "INSERT INTO recipes(title, description, cooking_time, ingredients, instructions) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	queryGetRecipe    = "SELECT id, title, description, cooking_time, ingredients, instructions, date_created FROM recipes WHERE id = $1;"
	queryListRecipes  = "SELECT id, title, description, cooking_time, ingredients, instructions, date_created FROM recipes;"
	queryUpdateRecipe = "UPDATE recipes SET title=$1, description=$2, cooking_time=$3, ingredients=$4, instructions=$5 WHERE id = $6;"
	queryDeleteRecipe = "DELETE FROM recipes WHERE id=$1;"
)

// GetRecipe returns an individual recipe by recipe id
func (recipe *Recipe) GetRecipe() *errors.RestErr {
	ingredients := strings.Join(recipe.Ingredients, ", ")

	stmt, err := recipes_db.Client.Prepare(queryGetRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare get recipe statement: %s", err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(recipe.Id)
	if err := result.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.CookingTime, &ingredients, &recipe.Instructions, &recipe.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf(
				"recipe id %d not found", recipe.Id))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get recipe %d: %s", recipe.Id, err.Error()))
	}

	return nil
}

// SaveRecipe adds a new recipe
func (recipe *Recipe) SaveRecipe() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryInsertRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}


	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(recipe.Title, &recipe.Description, &recipe.CookingTime, pq.StringArray(recipe.Ingredients), &recipe.Instructions).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueRecipeTitle) {
			return errors.NewBadRequestError(fmt.Sprintf(
				"recipe name %s already exists", recipe.Title))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save new recipe: %s", err.Error()))
	}

	recipe.Id = id

	return nil
}

// DeleteRecipe removes a recipe 
func (recipe *Recipe) DeleteRecipe(recipeId int64) *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryDeleteRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare delete recipe query: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipeId)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to delete recipe: %s", err.Error()))
	}
	return nil
}

// ListRecipes returns a set of recipes
func (recipe *Recipe) ListRecipes() ([]Recipe, *errors.RestErr) {
	stmt, err := recipes_db.Client.Prepare(queryListRecipes)
	if err != nil {
		return nil, errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare list recipes statement: %s", err.Error()))
	}
	defer stmt.Close()

	var recipes []Recipe
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(
			fmt.Sprintf("failed to list recipes: %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.CookingTime, &recipe.Ingredients, (*pq.StringArray(&recipe.Ingredients)), &recipe.DateCreated)
		if err != nil {
			return nil, errors.NewInternalServerError(
				fmt.Sprintf("there was a problem scanning rows for recipe list: %s", err.Error()))
		}
		recipes = append(recipes, *recipe)

		if len(recipes) == 0 {
			return nil, errors.NewNotFoundError("no recipes found")
		}
	}

	return recipes, nil
}

// UpdateRecipe updates an individual recipe
func (recipe *Recipe) UpdateRecipe() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryUpdateRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare update recipe statement: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipe.Title, recipe.Description, recipe.CookingTime, recipe.Ingredients, recipe.Instructions, recipe.Id)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to update recipe: %s", err.Error()))
	}

	return nil
}

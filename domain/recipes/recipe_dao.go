package recipes

import (
	"fmt"
	"github.com/dogray7722/golang-cookbook_dogray7722/datasources/postgres/recipes_db"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
	"strings"
)

const (
	indexUniqueRecipeName = "constraint_name"
	errorNoRows           = "no rows in result set"

	queryInsertRecipe = "INSERT INTO recipes(name, instructions, status) VALUES($1, $2, $3) RETURNING id;"
	queryGetRecipe    = "SELECT id, name, instructions, status, date_created FROM recipes WHERE id = $1"

	queryInsertIngredient       = "INSERT INTO ingredients(serving_size, item) VALUES($1, $2) RETURNING id;"
	queryGetIngredientsByRecipe = `SELECT id, serving_size, item FROM ingredients WHERE id IN (SELECT ingredient_id FROM recipes_to_ingredients WHERE recipe_id = $1)`

	queryInsertLookup = "INSERT INTO recipes_to_ingredients(recipe_id, ingredient_id) VALUES($1, $2);"
)

func (recipe *Recipe) Get() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryGetRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare get recipe statement: %s", err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(recipe.Id)
	if err := result.Scan(&recipe.Id, &recipe.Name, &recipe.Instructions, &recipe.Status, &recipe.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf(
				"recipe id %d not found", recipe.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get recipe %d: %s", recipe.Id, err.Error()))
	}

	recipeIngredients, err := getIngredients(recipe.Id)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("there was a problem retrieving recipe ingredients: %s", err.Error()))
	}

	recipe.Ingredients = recipeIngredients

	return nil

}

func getIngredients(recipeID int64) ([]Ingredient, error) {
	stmt, err := recipes_db.Client.Prepare(queryGetIngredientsByRecipe)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Ingredient{}
	for rows.Next() {
		ing := Ingredient{}
		err := rows.Scan(&ing.Id, &ing.ServingSize, &ing.Item)
		if err != nil {
			return nil, err
		}
		results = append(results, ing)
	}

	return results, nil
}

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
		//TODO Refactor to use error code
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

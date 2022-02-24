package recipes

import (
	"fmt"
	"github.com/golang-cookbook/datasources/postgres/recipes_db"
	"github.com/golang-cookbook/utils/errors"
	"strings"
)

const (
	indexUniqueRecipeName = "constraint_name"
	errorNoRows           = "no rows in result set"

	queryInsertRecipe = "INSERT INTO recipes(name, instructions, description, status) VALUES($1, $2, $3, $4) RETURNING id;"
	queryGetRecipe    = "SELECT id, name, instructions, description, status, date_created FROM recipes WHERE id = $1;"
	queryListRecipes  = "SELECT id, name, instructions, description, status, date_created FROM recipes;"
	queryUpdateRecipe = "UPDATE recipes SET name=$1, instructions=$2, description=$3, status=$4 WHERE id = $5;"
	queryDeleteRecipe = "DELETE FROM recipes WHERE id=$1;"

	queryInsertIngredient          = "INSERT INTO ingredients(serving_size, item) VALUES($1, $2) RETURNING id;"
	queryGetIngredientsByRecipe    = `SELECT id, serving_size, item FROM ingredients WHERE id IN (SELECT ingredient_id FROM recipes_to_ingredients WHERE recipe_id = $1);`
	queryDeleteIngredientsByRecipe = `DELETE FROM ingredients WHERE id IN (SELECT ingredient_id FROM recipes_to_ingredients WHERE recipe_id = $1);`

	queryInsertLookup = "INSERT INTO recipes_to_ingredients(recipe_id, ingredient_id) VALUES($1, $2);"
	queryDeleteLookup = "DELETE FROM recipes_to_ingredients WHERE recipe_id = $1;"
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
	err = stmt.QueryRow(recipe.Name, recipe.Instructions, recipe.Description, recipe.Status).Scan(&id)
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
		if err := recipe.saveRecipeIngredient(recipe.Id, id); err != nil {
			return errors.NewInternalServerError(
				fmt.Sprintf("failed to save recipe ingredients: %s", err))
		}
	}

	return nil
}

func (recipe *Recipe) saveRecipeIngredient(recipeId, ingredientId int64) *errors.RestErr {
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

func (recipe *Recipe) DeleteRecipe(recipeId int64) *errors.RestErr {
	if err := recipe.DeleteIngredients(recipeId); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to delete recipe ingredients"))
	}

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

func (recipe *Recipe) DeleteIngredients(recipeID int64) *errors.RestErr {

	stmt, err := recipes_db.Client.Prepare(queryDeleteIngredientsByRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare delete recipe ingredients query: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipeID)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to delete ingredients by recipe: %s", err.Error()))
	}

	if err := recipe.deleteRecipeIngredient(recipeID); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to delete recipe ingredient relationships: %s", err))
	}

	return nil

}

func (recipe *Recipe) deleteRecipeIngredient(recipeId int64) *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryDeleteLookup)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare recipe to ingrdient delete statement: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipeId)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to delete recipe to ingredients relationship: %s", err.Error()))
	}

	return nil
}

func (recipe *Recipe) List() ([]Recipe, *errors.RestErr) {
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
		err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Instructions, &recipe.Description, &recipe.Status, &recipe.DateCreated)
		if err != nil {
			return nil, errors.NewInternalServerError(
				fmt.Sprintf("there was a problem scanning rows for recipe list: %s", err.Error()))
		}

		recipeIngredients, err := getIngredients(recipe.Id)
		if err != nil {
			return nil, errors.NewInternalServerError(
				fmt.Sprintf("there was a problem retrieving recipe ingredients: %s", err.Error()))
		}

		recipe.Ingredients = recipeIngredients

		recipes = append(recipes, *recipe)

		if len(recipes) == 0 {
			return nil, errors.NewNotFoundError("no recipes found")
		}
	}

	return recipes, nil
}

func (recipe *Recipe) Update() *errors.RestErr {
	stmt, err := recipes_db.Client.Prepare(queryUpdateRecipe)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to prepare update recipe statement: %s", err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipe.Name, recipe.Instructions, recipe.Description, recipe.Status, recipe.Id)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to update recipe: %s", err.Error()))
	}

	if err := recipe.SaveIngredients(); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("failed to save ingredeints: %s", err))
	}

	return nil
}

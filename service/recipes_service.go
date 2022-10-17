package service

import (
	"context"

	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
	"github.com/golang-cookbook/utils/errors"
)

func CreateRecipe(recipe db.CreateRecipeParams) (*db.Recipe, *errors.RestErr) {
	// need instance of queries in which to run createRecipe
	// needs to be global
	
	
	res, err := queries.CreateRecipe(context.Background(), recipe)
	if err != nil {
		return nil, &errors.RestErr{Error: err.Error()}
	}

	return &res, nil
	
}

func GetRecipe(recipeId int64) (*recipes_db.Recipe, *errors.RestErr) {
	if recipeId <= 0 {
		return nil, errors.NewBadRequestError("invalid recipe id")
	}
	result := &recipes_db.Recipe{Id: recipeId}
	if err := result.GetRecipe(); err != nil {
		return nil, err
	}
	return result, nil
}

func ListRecipes() ([]recipes.Recipe, *errors.RestErr) {
	dao := &recipes.Recipe{}
	recipes, err := dao.ListRecipes()
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func UpdateRecipe(recipe recipes.Recipe) (*recipes.Recipe, *errors.RestErr) {
	current, err := GetRecipe(recipe.Id)
	if err != nil {
		return nil, err
	}

	current.Title = recipe.Title
	current.Description = recipe.Description
	current.CookingTime = recipe.CookingTime
	current.Ingredients = recipe.Ingredients
	current.Instructions = recipe.Instructions

	err = current.UpdateRecipe()
	if err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteRecipe(recipeId int64) *errors.RestErr {
	if recipeId <= 0 {
		return errors.NewBadRequestError("invalid recipe id")
	}
	recipe := &recipes.Recipe{
		Id: recipeId,
	}
	return recipe.DeleteRecipe(recipeId)
}

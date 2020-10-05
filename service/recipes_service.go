package service

import (
	"github.com/dogray7722/golang-cookbook_dogray7722/domain/recipes"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
)

func CreateRecipe(recipe recipes.Recipe) (*recipes.Recipe, *errors.RestErr) {
	if err := recipe.Save(); err != nil {
		return nil, err
	}

	if err := recipe.SaveIngredients(); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func GetRecipe(recipeId int64) (*recipes.Recipe, *errors.RestErr) {
	if recipeId <= 0 {
		return nil, errors.NewBadRequestError("invalid recipe id")
	}
	result := &recipes.Recipe{Id: recipeId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func ListRecipes() (*recipes.Recipe, *errors.RestErr) {
	results := &recipes.Recipe{}
	if err := results.List(); err != nil {
		return nil, err
	}
	return results, nil
}

func UpdateRecipe(recipe recipes.Recipe) (*recipes.Recipe, *errors.RestErr) {
	current, err := GetRecipe(recipe.Id)
	if err != nil {
		return nil, err
	}

	err = current.DeleteIngredients(recipe.Id)
	if err != nil {
		return nil, err
	}

	current.Name = recipe.Name
	current.Instructions = recipe.Instructions
	current.Ingredients = recipe.Ingredients
	current.Status = recipe.Status

	err = current.Update()
	if err != nil {
		return nil, err
	}

	return current, nil
}

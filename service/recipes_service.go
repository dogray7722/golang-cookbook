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

//func GetRecipe(recipeId int64) (*recipes.Recipe, *errors.RestErr) {
//	if recipeId <= 0 {
//		return nil, errors.NewBadRequestError("invalid recipe id")
//	}
//	result := &recipes.Recipe{Id: recipeId}
//	if err := result.Get(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}

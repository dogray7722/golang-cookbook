package service

import (
	"github.com/dogray7722/golang-cookbook/domain/recipes"
	"github.com/dogray7722/golang-cookbook/utils/errors"
)

func CreateRecipe(recipe recipes.Recipe) (*recipes.Recipe, *errors.RestErr) {
	if err := recipe.Save(); err != nil {
		return nil, err
	}

	return &recipe, nil

}

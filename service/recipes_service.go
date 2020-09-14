package service

import "github.com/dogray7722/golang-cookbook/domain/recipes"

func CreateRecipe(recipe recipes.Recipe) (*recipes.Recipe, error) {
	return &recipe, nil
}
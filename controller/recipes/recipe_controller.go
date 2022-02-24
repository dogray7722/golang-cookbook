package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-cookbook/service"
	"github.com/golang-cookbook/domain/recipes"
	"github.com/golang-cookbook/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := service.CreateRecipe(recipe)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func List(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	recipes, listErr := service.ListRecipes()
	if listErr != nil {
		c.JSON(listErr.Status, listErr)
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func Get(c *gin.Context) {
	recipeId, recipeErr := strconv.ParseInt(c.Param("recipe_id"), 10, 64)
	if recipeErr != nil {
		err := errors.NewBadRequestError("invalid recipe id")
		c.JSON(err.Status, err)
		return
	}
	recipe, getErr := service.GetRecipe(recipeId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func Update(c *gin.Context) {
	recipeId, recipeErr := strconv.ParseInt(c.Param("recipe_id"), 10, 64)
	if recipeErr != nil {
		err := errors.NewBadRequestError("recipe id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	recipe.Id = recipeId

	result, err := service.UpdateRecipe(recipe)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	recipeId, recipeErr := strconv.ParseInt(c.Param("recipe_id"), 10, 64)
	if recipeErr != nil {
		err := errors.NewBadRequestError("invalid recipe id")
		c.JSON(err.Status, err)
		return
	}

	if deleteErr := service.DeleteRecipe(recipeId); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

package app

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-cookbook/utils/errors"

	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
)

type createRecipeRequest struct {
	Title        string         `json:"title" binding:"required"`
	Description  sql.NullString `json:"description"`
	CookingTime  string         `json:"cookingTime" binding:"required"`
	Ingredients  []string       `json:"ingredients" binding:"required"`
	Instructions string         `json:"instructions" binding:"required"`
}

func (s *Server) createRecipe(ctx *gin.Context) {
	var req createRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	arg := db.CreateRecipeParams{
		Title: req.Title,
		Description: req.Description,
		CookingTime: req.CookingTime,
		Ingredients: req.Ingredients,
		Instructions: req.Instructions,
	}

	result, saveErr := s.store.createRecipe

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func List(c *gin.Context) {
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

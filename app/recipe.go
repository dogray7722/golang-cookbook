package app

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
)

type createRecipeRequest struct {
	Title string `json:"title" binding:"required"`
	Description  string `json:"description"`
	CookingTime  string         `json:"cookingTime" binding:"required"`
	Ingredients  []string       `json:"ingredients" binding:"required"`
	Instructions string         `json:"instructions" binding:"required"`
}

type getRecipeRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type updateRecipeRequest struct {
	ID string `json:"id"`
	Title string `json:"title" binding:"required"`
	Description  string `json:"description"`
	CookingTime  string         `json:"cookingTime" binding:"required"`
	Ingredients  []string       `json:"ingredients" binding:"required"`
	Instructions string         `json:"instructions" binding:"required"`
}

func (server *Server) createRecipe(ctx *gin.Context) {
	var req createRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRecipeParams{
		Title: req.Title,
		Description:  req.Description,
		CookingTime:  req.CookingTime,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
	}

	result, err := server.store.CreateRecipe(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) getRecipe(ctx *gin.Context) {
	var req getRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	recipe, err := server.store.GetRecipe(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}

func (server *Server) listRecipes(ctx *gin.Context) {
	var params db.ListRecipesParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	recipes, err := server.store.ListRecipes(ctx, params) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	
	ctx.JSON(http.StatusOK, recipes)
}

func (server *Server) updateRecipe(ctx *gin.Context) {
	var req updateRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	recipeId, err := strconv.Atoi(ctx.Param("recipe_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateRecipeParams{
		ID: int32(recipeId),
		Title: req.Title,
		Description:  req.Description,
		CookingTime:  req.CookingTime,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
	}

	result, err := server.store.UpdateRecipe(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) deleteRecipe(ctx *gin.Context) {
	recipeId, err := strconv.Atoi(ctx.Param("recipe_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteRecipe(ctx, int32(recipeId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}


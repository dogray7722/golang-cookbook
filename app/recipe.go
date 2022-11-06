package app

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
)

// createRecipeRequest defines the recipe create request object
type createRecipeRequest struct {
	Title string `json:"title" binding:"required"`
	Description  string `json:"description"`
	CookingTime  string         `json:"cookingTime" binding:"required"`
	Ingredients  []string       `json:"ingredients" binding:"required"`
	Instructions string         `json:"instructions" binding:"required"`
}

// getRecipeRequest defines the recipe get request object
type getRecipeRequest struct {
	ID int32 `uri:"recipe_id" binding:"required,min=1"`
}

// updateRecipeRequest defines the recipe update request object
type updateRecipeRequest struct {
	ID string `json:"id"`
	Title string `json:"title" binding:"required"`
	Description  string `json:"description"`
	CookingTime  string         `json:"cookingTime" binding:"required"`
	Ingredients  []string       `json:"ingredients" binding:"required"`
	Instructions string         `json:"instructions" binding:"required"`
}

// listRecipesRequest defines the recipe list request object
type listRecipesRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// createRecipe is the api handler for creating a single recipe
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

// getRecipe is the api handler for retrieving a single recipe
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

// listRecipes is the api handler for retrieving a list of recipes
func (server *Server) listRecipes(ctx *gin.Context) {
	var req listRecipesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRecipesParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	
	recipes, err := server.store.ListRecipes(ctx, arg) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	
	ctx.JSON(http.StatusOK, recipes)
}

// updateRecipe is the api handler for modifying a single recipe
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

// deleteRecipe is the api handler for deleting a single recipe
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


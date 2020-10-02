package recipes

import (
	"github.com/dogray7722/golang-cookbook_dogray7722/domain/recipes"
	"github.com/dogray7722/golang-cookbook_dogray7722/service"
	"github.com/dogray7722/golang-cookbook_dogray7722/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
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

}

//func Get(c *gin.Context) {
//	recipeId, recipeErr := strconv.ParseInt(c.Param("recipe_id"), 10, 64)
//	if recipeErr != nil {
//		err := errors.NewBadRequestError("invalid recipe id")
//		c.JSON(err.Status, err)
//		return
//	}
//	recipe, getErr := service.GetRecipe(recipeId)
//	if getErr != nil {
//		c.JSON(getErr.Status, getErr)
//		return
//	}
//	c.JSON(http.StatusOK, recipe)
//}

func Update(c *gin.Context) {
	//TODO Implement
}

func Delete(c *gin.Context) {
	//TODO Implement
}

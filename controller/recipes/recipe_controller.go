package recipes

import (
	"github.com/dogray7722/golang-cookbook/domain/recipes"
	"github.com/dogray7722/golang-cookbook/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		//TODO return bad request to the caller
		return
	}


	result, saveErr := service.CreateRecipe(recipe)
	if saveErr != nil {
		//TODO: handle recipe creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func List(c *gin.Context) {
	//TODO Implement
}

func Get(c *gin.Context) {
	//TODO Implement
}

func Update(c *gin.Context) {
	//TODO Implement
}

func Delete(c *gin.Context) {
	//TODO Implement
}
package app

import (
	"github.com/dogray7722/golang-cookbook_dogray7722/controller/ping"
	"github.com/dogray7722/golang-cookbook_dogray7722/controller/recipes"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/recipes", recipes.Create)
	router.GET("/recipes/:recipe_id", recipes.Get)
	router.GET("/recipes", recipes.List)
	router.PUT("/recipes:recipe_id", recipes.Update)
	router.DELETE("/recipes/:recipe_id", recipes.Delete)
}

package app

import (
	"golang-cookbook/controller/ping"
	"golang-cookbook/controller/recipes"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/recipes", recipes.Create)
	router.GET("/recipes/:recipe_id", recipes.Get)
	router.GET("/recipes", recipes.List)
	router.PUT("/recipes/:recipe_id", recipes.Update)
	router.DELETE("/recipes/:recipe_id", recipes.Delete)
}

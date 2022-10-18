package recipes_db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang-cookbook/util"
	"github.com/stretchr/testify/require"
)

func PopulateTest() CreateRecipeParams {
	arg := CreateRecipeParams{
		Title: util.RandomTitle(),
		Description: sql.NullString{String: util.RandomDescription(), Valid: true},
		CookingTime: util.RandomCookingTime(),
		Ingredients: util.RandomIngredients(),
		Instructions: util.RandomInstructions(),
		DateCreated: time.Now().UTC(),
	}
	
	return arg
}

func createTestRecipe(t *testing.T) Recipe {
	arg := PopulateTest()

	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)

	require.Equal(t, arg.Title, recipe.Title)
	require.Equal(t, arg.Description, recipe.Description)
	require.Equal(t, arg.CookingTime, recipe.CookingTime)
	require.ElementsMatch(t, arg.Ingredients, recipe.Ingredients)
	require.Equal(t, arg.Instructions, recipe.Instructions)

	require.NotZero(t, recipe.ID)
	require.NotZero(t, recipe.DateCreated)

	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createTestRecipe(t)
} 

func TestGetRecipe(t *testing.T) {
	recipe := createTestRecipe(t)
	res, err := testQueries.GetRecipe(context.Background(), recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, recipe.ID, res.ID)
	require.Equal(t, recipe.Description, res.Description)
	require.Equal(t, recipe.CookingTime, res.CookingTime)
	require.Equal(t, recipe.Instructions, res.Instructions)
	require.ElementsMatch(t, recipe.Ingredients, res.Ingredients)
	
	require.NotZero(t, res.DateCreated)
}

func TestUpdateRecipe(t *testing.T) {
	recipe1 := createTestRecipe(t)

	originalRecipe, _ := testQueries.GetRecipe(context.Background(), recipe1.ID)
	arg := UpdateRecipeParams{
		ID: originalRecipe.ID,
		Title: util.RandomTitle(),
		Description: originalRecipe.Description,
		CookingTime: originalRecipe.CookingTime,
		Ingredients: originalRecipe.Ingredients,
		Instructions: originalRecipe.Instructions,
	}


	recipe2, err := testQueries.UpdateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe2)
	require.Equal(t, arg.Title, recipe2.Title)
	require.Equal(t, recipe1.Description, recipe2.Description)
	require.Equal(t, recipe1.CookingTime, recipe2.CookingTime)
	require.Equal(t, recipe1.Instructions, recipe2.Instructions)
	require.Equal(t, recipe1.Ingredients, recipe2.Ingredients)
	require.WithinDuration(t, recipe1.DateCreated, recipe2.DateCreated, time.Second)

}



// func TestListRecipes(t *testing.T) {
// 	recipe1 := createTestRecipe(t)
// 	recipe2 := createTestRecipe(t)
	
// 	res, err := testQueries.ListRecipes(context.Background())
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res)
// 	require.NotEqual(t, res[0].ID, res[len(res)-1].ID)
	
// }
package recipes_db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateRecipe(t *testing.T) {
	arg := CreateRecipeParams{
		Title: "PB&J",
		Description: sql.NullString{String: "peanut butter and jelly sandwich", Valid: true},
		CookingTime: "5 minutes",
		Ingredients: []string{"peanut butter", "jelly", "bread"},
		Instructions: "Spread the peanut butter and jelly on two pieces of bread, respectively. Stick the two gooey sides together",
		DateCreated: time.Now().UTC(),
	}

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
}
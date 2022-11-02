package app

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/golang-cookbook/datasources/postgres/recipes_db/mock"
	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
	"github.com/golang-cookbook/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestGetRecipeAPI runs mock tests against
func TestGetRecipeAPI(t *testing.T) {
	recipe := randomRecipe()

	testCases := []struct{
		name string
		recipeID int32
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			recipeID: recipe.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetRecipe(gomock.Any(), gomock.Eq(recipe.ID)).
				Times(1).
				Return(recipe, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchRecipe(t, recorder.Body, recipe)
			},
		},
		{
			name: "NotFound",
			recipeID: recipe.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetRecipe(gomock.Any(), gomock.Eq(recipe.ID)).
				Times(1).
				Return(db.Recipe{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			recipeID: recipe.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetRecipe(gomock.Any(), gomock.Eq(recipe.ID)).
				Times(1).
				Return(db.Recipe{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			recipeID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetRecipe(gomock.Any(), gomock.Any()).
				Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}


	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
	
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			
			server := NewServer(store)
			recorder := httptest.NewRecorder()
	
			url := fmt.Sprintf("/recipes/%d", tc.recipeID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
	
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomRecipe() db.Recipe {
	return db.Recipe{
		ID: int32(util.RandomInt(1, 1000)),
		Title: util.RandomTitle(),
		Description: util.RandomDescription(),
		CookingTime: util.RandomCookingTime(),
		Ingredients: util.RandomIngredients(),
		Instructions: util.RandomInstructions(),
	}
}

// requireBodyMatchRecipe checks whether the request body is equal to the recipe defined in the test
func requireBodyMatchRecipe(t *testing.T, body *bytes.Buffer, recipe db.Recipe) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotRecipe db.Recipe
	err = json.Unmarshal(data, &gotRecipe)
	require.NoError(t, err)
	require.Equal(t, recipe, gotRecipe)
}
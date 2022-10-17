// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe.sql

package recipes_db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipes (
  title,
  description,
  cooking_time,
  ingredients,
  instructions,
  date_created
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, title, description, cooking_time, ingredients, instructions, date_created
`

type CreateRecipeParams struct {
	Title        string         `json:"title"`
	Description  sql.NullString `json:"description"`
	CookingTime  string         `json:"cookingTime"`
	Ingredients  []string       `json:"ingredients"`
	Instructions string         `json:"instructions"`
	DateCreated  time.Time      `json:"dateCreated"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.Title,
		arg.Description,
		arg.CookingTime,
		pq.Array(arg.Ingredients),
		arg.Instructions,
		arg.DateCreated,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CookingTime,
		pq.Array(&i.Ingredients),
		&i.Instructions,
		&i.DateCreated,
	)
	return i, err
}

const listRecipes = `-- name: ListRecipes :many
SELECT id, title, description, cooking_time, ingredients, instructions, date_created FROM recipes
ORDER BY title
`

func (q *Queries) ListRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CookingTime,
			pq.Array(&i.Ingredients),
			&i.Instructions,
			&i.DateCreated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

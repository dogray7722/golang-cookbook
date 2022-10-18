-- name: CreateRecipe :one
INSERT INTO recipes (
  title,
  description,
  cooking_time,
  ingredients,
  instructions,
  date_created
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetRecipe :one
  SELECT * FROM recipes
  WHERE id = $1 LIMIT 1;

-- name: ListRecipes :many
SELECT * FROM recipes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;

-- name: UpdateRecipe :one
UPDATE recipes
  set title = $2,
  description = $3,
  cooking_time = $4,
  ingredients = $5,
  instructions = $6
WHERE id = $1
RETURNING *;

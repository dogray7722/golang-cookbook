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
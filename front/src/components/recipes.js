import React, { useEffect, useState }from 'react'
import { Card, CardHeader, Typography, makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  recipeCard: {
    margin: "10px",
    padding: "10px"
  }
})

export default function Recipes() {
  const [recipes, setRecipes] = useState([])

  useEffect(() => {
    loadData();
  }, [])

  const classes = useStyles()

  const loadData = async () => {
    const response = await fetch("http://localhost:8080/recipes")
    const data = await response.json();
    setRecipes(data)
  }

const content = recipes.map(recipe => {
  return (
  <Card
      raised={true}
      key={recipe.id}
      className={classes.recipeCard}
  >
      <Typography
          variant="h6"
          color="textSecondary"
      >
        {recipe.name}
      </Typography>

    <ul>
      {recipe.ingredients.map(ingredient => (
      <li key={ingredient.id}>
        {`${ingredient.serving_size} ${ingredient.item}`}      
      </li>
      ))} 
    </ul>
    <p>{recipe.instructions}</p>
  </Card>
  );
})

  return (
    <>
      {content}
    </>
  )
}


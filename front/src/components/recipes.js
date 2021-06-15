import React, { useEffect, useState }from 'react'
import { Box, Typography } from '@material-ui/core'

const Recipes = () => {
  const [recipes, setRecipes] = useState([])

  useEffect(() => {
    loadData();
  }, [])

  const loadData = async () => {
    const response = await fetch("http://localhost:8080/recipes")
    const data = await response.json();
    setRecipes(data)
  }

const content = recipes.map(recipe => {
  return (
  <Box component="div" key={recipe.id}>

    <Typography
        variant="h6"
        gutterBottom
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
  </Box>
  );
})

  return (
    <>
      <Typography
          variant="h3"
          align="center"
          color="textSecondary"
          gutterBottom
      >
        Recipes
      </Typography>
      {content}
    </>
  )
}

export default Recipes

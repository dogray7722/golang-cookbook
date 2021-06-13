import React, { useEffect, useState }from 'react'
import { Box } from '@material-ui/core'

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

    <h3>{recipe.name}</h3>
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
      <h2>Recipes</h2>
      {content}
    </>
  )
}

export default Recipes

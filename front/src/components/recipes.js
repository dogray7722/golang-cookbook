import React, { useEffect, useState }from 'react'

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
  <div key={recipe.id}>
    <h3>{recipe.name}</h3>
    <ul>
      {recipe.ingredients.map(ingredient => (
      <li key={ingredient.id}>
        {`${ingredient.serving_size} ${ingredient.item}`}      
      </li>
      ))} 
    </ul>
    <p>{recipe.instructions}</p>
    </div>
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

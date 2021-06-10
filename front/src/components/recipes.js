import React, { useEffect, useState }from 'react'

const Recipes = () => {
  
  // const [recipes, setRecipes] = useState([])

  useEffect(() => {
    loadData();
  }, [])

  const loadData = async () => {
    const response = await fetch("http://localhost:8080/recipes")
    const data = await response.json();
    console.log(data)
  }


  return (
    <>
      <h2>Recipes</h2>
    </>
  )
}

export default Recipes

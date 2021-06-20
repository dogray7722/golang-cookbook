import React, { useEffect, useState }from 'react'
import {Card, CardHeader, makeStyles, Avatar, CardContent, Typography} from '@material-ui/core'
import purple from "@material-ui/core/colors/purple";

const useStyles = makeStyles({
  recipeCard: {
    margin: "10px",
    padding: "10px"
  },
  avatar: {
    backgroundColor: purple[500]
  },
  content: {
    padding: "8px"
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
    <CardHeader
      titleTypographyProps={{variant:'h6' }}
      title={recipe.name}
      avatar={
        <Avatar
          aria-label="recipe" className={classes.avatar}
        >
          {recipe.name.charAt(0).toUpperCase()}
        </Avatar>
      }
    />
    <CardContent className={classes.content}>
      <Typography variant="subtitle1" color="textPrimary">
        Ingredients
      </Typography>
      <Typography color="textSecondary">
      <ul style={{listStyle: "none", margin: "3px", fontSize: "14px", fontFamily: "Roboto, Helvetica, Arial, sans-serif"}}>
        {recipe.ingredients.map(ingredient => (
            <li key={ingredient.id}>
              {`${ingredient.serving_size} ${ingredient.item}`}
            </li>
        ))}
      </ul>
      </Typography>
      <Typography variant="subtitle1" color="textPrimary">
        Instructions
      </Typography>
      <Typography variant="body2" color="textSecondary" component="p">
        {recipe.instructions}
      </Typography>
    </CardContent>
  </Card>
  );
})

  return (
    <>
      {content}
    </>
  )
}


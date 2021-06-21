import React, { useEffect, useState }from 'react'
import {Avatar, Card, CardActions, CardContent, CardHeader, Collapse, IconButton, makeStyles, Typography } from '@material-ui/core'
import purple from "@material-ui/core/colors/purple";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import clsx from 'clsx';

const useStyles = makeStyles((theme) => ({
  recipeCard: {
    maxWidth: 350,
    margin: "10px",
    padding: "10px"
  },
  avatar: {
    backgroundColor: purple[500]
  },
  content: {
    padding: "8px"
  },
  expand: {
    transform: 'rotate(0deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    })
  },
  expandOpen: {
    transform: 'rotate(180deg)'
  }
}))

export default function Recipes() {
  const classes = useStyles()
  const [recipes, setRecipes] = useState([])
  const [expandedId, setExpandedId] = useState(-1)

  useEffect(() => {
    loadData();
  }, [])


  const handleExpandClick = i => {
    setExpandedId(expandedId === i ? -1 : i)
  }

  const loadData = async () => {
    const response = await fetch("http://localhost:8080/recipes")
    const data = await response.json();
    setRecipes(data)
  }

const content = recipes.map((recipe, i) => {
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
      <Typography variant="body1" color="textPrimary" >
        {recipe.description}
      </Typography>
    </CardContent>
    <CardActions>
      <IconButton
        className={clsx(classes.expand, {
          [classes.expandOpen]: expandedId
        })}
        onClick={() => handleExpandClick(i)}
        aria-expanded={expandedId === i}
        aria-label="show more"
      >
        <ExpandMoreIcon />
      </IconButton>
    </CardActions>
    <Collapse in={expandedId === i} timeout="auto" unmountOnExit>
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
    </Collapse>
  </Card>
  );
})

  return (
    <>
      {content}
    </>
  )
}


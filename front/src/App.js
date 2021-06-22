import React from 'react';
import { Container, Typography } from '@material-ui/core';
import Recipes from './components/recipes'
import RecipeForm from "./components/recipeForm";


function App() {

  return (
    <>
    <Container maxWidth="sm">
        <Typography
            color="secondary"
            variant="h2"
            align="center"
        >
            Favorite Recipes
        </Typography>
        <RecipeForm />
        <Recipes />
    </Container>
    
    </>
  
  );
}

export default App;

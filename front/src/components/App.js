import React from 'react';
import { Container, Typography } from '@material-ui/core';
import Recipes from './recipes'
import RecipeForm from "./recipeForm";
import Header from '../components/ui/header'


function App() {

  return (
    <>
    <Container maxWidth="sm">
        <Header/>
        <Typography
            color="secondary"
            variant="h2"
            align="center"
        >
            Favorite Recipes
        </Typography>
        {/*<RecipeForm />*/}
        {/*<Recipes />*/}
    </Container>
    
    </>
  
  );
}

export default App;

import React from 'react';
import { ThemeProvider } from "@material-ui/styles";
import theme from './ui/theme';
import { Container, Typography } from '@material-ui/core';
import Recipes from './ui/recipes'
import RecipeForm from "./ui/recipeForm";
import Header from '../components/ui/header'


function App() {

  return (
    <ThemeProvider theme={theme}>
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
    </ThemeProvider>
  );
}

export default App;

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
        {[...new Array(120)]
            .map(
                () => `Cras mattis consectetur purus sit amet fermentum.
                Cras justo odio, dapibus ac facilisis in, egestas eget quam.
                Morbi leo risus, porta ac consectetur ac, vestibulum at eros.
                Praesent commodo cursus magna, vel scelerisque nisl consectetur et.`,
            )
            .join('\n')}
        {/*<RecipeForm />*/}
        {/*<Recipes />*/}
    </Container>
    
    </>
  
  );
}

export default App;

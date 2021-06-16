import React from 'react';
import { Container, Typography } from '@material-ui/core';
import Recipes from './components/recipes'

function App() {

  return (
    <>
    <Container maxWidth="sm">
        <Typography
            color="primary"
            variant="h2"
            align="center"
            gutterBottom
        >
            Favorite Recipes
        </Typography>
        <Recipes />
    </Container>
    
    </>
  
  );
}

export default App;

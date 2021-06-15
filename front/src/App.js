import React from 'react';
import { Container, Paper, Typography } from '@material-ui/core';
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
            Super Dope Cookbook
        </Typography>
      <Paper style={{padding: "20px"}} elevation={5} variant="filled">
        <Recipes />
      </Paper>
    </Container>
    
    </>
  
  );
}

export default App;

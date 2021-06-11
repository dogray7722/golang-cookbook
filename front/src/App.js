import React from 'react';
import { Container, Paper } from '@material-ui/core';
import Recipes from './components/recipes'

function App() {

  return (
    <Container maxWidth="sm">
      <h1>Go Lang Cookbook</h1>
      <Paper style={{padding: "20px"}} elevation={5} variant="outlined">
        <Recipes />
      </Paper>
    </Container>
  );
}

export default App;

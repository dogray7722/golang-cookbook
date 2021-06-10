import React from 'react';
import { Container } from '@material-ui/core';
import Recipes from './components/recipes'

function App() {

  return (
    <Container maxWidth="sm">
      <h1>Go Lang Cookbook</h1>
      <Recipes />
    </Container>
  );
}

export default App;

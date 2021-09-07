import React, { useState } from 'react'
import {
    Button,
    Grid,
    TextareaAutosize,
    TextField
} from '@material-ui/core'

export default function RecipeForm() {
    const [open, setOpen] = useState(false)

    const handleClickOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }


    return(
        <>
            <Grid container s={4} justify="center">
                <Button color='primary' variant='contained' size='large' onClick={handleClickOpen}>Create New Recipe</Button>
            </Grid>
            <form>
               <TextareaAutosize rowsMin={5} placeholder="Enter your recipe here">

               </TextareaAutosize>
            </form>
        </>
    )
}


import React, { useState } from 'react'
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
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
            <Button color='primary' variant='contained' size='large' onClick={handleClickOpen}>Create New Recipe</Button>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Create Recipe</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Recipe Description
                    </DialogContentText>
                    <TextareaAutosize
                        autoComplete={false}
                        id="description"
                        label="Description"
                        type="text"
                        required={true}
                        rowsMin={3}

                    />
                </DialogContent>
                <DialogContent>
                <DialogContentText>
                    Recipe Instructions
                </DialogContentText>
                <TextareaAutosize
                    autoComplete={false}
                    id="instructions"
                    label="Instructions"
                    type="text"
                    required={true}
                    rowsMin={3}

                />
            </DialogContent>
                <DialogContent>
                    <TextField
                        autoComplete={false}
                        id="ingredients"
                        label="Ingredients"
                        type="textArea"
                        required={true}
                        variant="outlined"

                    />
                </DialogContent>
                <DialogActions>
                </DialogActions>
            </Dialog>
        </>
    )
}


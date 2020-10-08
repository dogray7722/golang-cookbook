# golang-cookbook

## Purpose
This application demonstrates knowledge of creating a REST-Full Web API using GO.  It incorporates the gin/gonic web server package, and uses the pq postgres driver for the database/sql instance.

## Instructions
- Run the migration script in datasources/postgres/recipes_db/db_migrations/initial_migration_up.sql
- Run the application locally (AWS instance is forthcoming)

## Testing CRUD Operations

#### Create
Call the create endpoint at `localhost:8080/recipes`
Send a json request in the following format:
```
{
    "name": "Albondigas",
    "ingredients": [
        {
            "serving_size": "1 pound",
            "item": "ground buffalo meat"
        },
        {
            "serving_size": "2",
            "item": "eggs"
        },
        {
            "serving_size": "1 tablespoon",
            "item": "vegetable oil"
        },
        {
            "serving_size": "1 can",
            "item": "fire roasted diced tomatoes"
        },
        {
            "serving_size": "1 cup",
            "item": "cooked rice"
        }
    ],
    "instructions": "Parboil the eggs.  Shell the eggs.  Brown the meat.  Make meatballs around the eggs.  
    Add salt, pepper, and spices to your liking.  Simmer with the fire roasted tomatoes 20 minutes or until cooked through.  
    Serve over the rice.",
    "status": "active"
}
```

### Get
Use the resultant recipe ID from the create call for the GET endpoint: `localhost:8080/recipes/{recipeId}`

### Update
Modify the JSON in the create example above and send a PUT request to the same endpoint as the GET request

### Delete 
Make a DELETE request to the same endpoint as GET / UPDATE

### List
Make a GET request to `localhost:8080/recipes`

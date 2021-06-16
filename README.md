# golang-cookbook

## Purpose
This application demonstrates full stack chops including a REST-Full Web API using GO, React Hooks, and Material UI for styling.  It incorporates the gin/gonic web server package, and uses the pq postgres driver for the database/sql instance.

## Setup
- Configure the following local environment variables to set up your database connection:
    - postgres_host
    - postgres_port
    - postgres_user
    - postgres_password
    - postgres_dbname 
- Create and connect to the database and run the migration script in datasources/postgres/recipes_db/db_migrations/initial_migration_up.sql
- Run the application locally
- Navigate to the front directory and run `npm start`

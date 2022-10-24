package main

import (
	"database/sql"
	"log"


	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
) 

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	store := db.NewStore(conn)

	server := 
	


}

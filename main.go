package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/golang-cookbook/app"
	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	store := db.NewStore(conn)
	server := app.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

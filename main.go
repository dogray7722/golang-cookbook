package main

import (
	"database/sql"
	"log"

	"github.com/golang-cookbook/app"
	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
	"github.com/golang-cookbook/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}	

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	store := db.NewStore(conn)
	server := app.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

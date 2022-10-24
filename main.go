package main

import "github.com/golang-cookbook/app"

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable"
)

func main() {
	app.StartApplication()
}

package recipes_db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/recipes_db?sslmode=disable"
)

var testQueries Store

func TestMain(m *testing.M) {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

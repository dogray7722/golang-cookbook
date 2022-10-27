package recipes_db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-cookbook/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../../") 
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

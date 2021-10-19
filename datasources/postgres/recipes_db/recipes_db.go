package recipes_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgresHost  = "postgres_host"
	postgresPort      = "postgres_port"
	postgresUser     = "postgres_user"
	postgresPassword = "postgres_password"
	postgresDbname   = "postgres_dbname"
)

var (
	Client   *sql.DB
	host     = os.Getenv(postgresHost)
	port     = os.Getenv(postgresPort)
	user     = os.Getenv(postgresUser)
	password = os.Getenv(postgresPassword)
	dbname   = os.Getenv(postgresDbname)
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)

	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

package recipes_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "postgres_host"
	postgres_port     = "postgres_port"
	postgres_user     = "postgres_user"
	postgres_password = "postgres_password"
	postgres_dbname   = "postgres_dbname"
)

var (
	Client   *sql.DB
	host     = os.Getenv(postgres_host)
	port     = os.Getenv(postgres_port)
	user     = os.Getenv(postgres_user)
	password = os.Getenv(postgres_password)
	dbname   = os.Getenv(postgres_dbname)
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
	//defer Client.Close()

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

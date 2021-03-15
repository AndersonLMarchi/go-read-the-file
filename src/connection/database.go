package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Vars to get environment variables
var Vars map[string]string

func Connect() *sql.DB {

	Vars, _ = godotenv.Read("../../.env")

	// pgsql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Vars["POSTGRES_HOST"], Vars["POSTGRES_PORT"], Vars["POSTGRES_USER"], Vars["POSTGRES_PASSWORD"], Vars["POSTGRES_DB"])
	pgsql := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", Vars["POSTGRES_USER"], Vars["POSTGRES_PASSWORD"], Vars["POSTGRES_HOST"], Vars["POSTGRES_PORT"], Vars["POSTGRES_DB"])
	result, err := sql.Open("postgres", pgsql)
	if err != nil {
		log.Fatalf("Something wrong is not correct : %s", err)
	}

	return result

}

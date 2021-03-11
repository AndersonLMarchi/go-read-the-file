package connection

import (
	"database/sql"
	"fmt"
	"log"
)

func connection() *sql.DB {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "123456"
	dbname := "reader_db"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	return result

}

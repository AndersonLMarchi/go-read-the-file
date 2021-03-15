// package main

// import (
// 	"fmt"
// 	connection "go-reader/src/connection"
// )

// var db = connection.Connect()

// func main() {

// 	// var Vars map[string]string
// 	// Vars, _ = godotenv.Read(".env")
// 	// fmt.Println(vars["POSTGRES_DB"])

// 	// pgsql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Vars["POSTGRES_HOST"], Vars["POSTGRES_PORT"], Vars["POSTGRES_USER"], Vars["POSTGRES_PASSWORD"], Vars["POSTGRES_DB"])
// 	// pgsql := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full", Vars["POSTGRES_USER"], Vars["POSTGRES_PASSWORD"], Vars["POSTGRES_HOST"], Vars["POSTGRES_PORT"], Vars["POSTGRES_DB"])

// 	sql := "select * from information_schema.columns"

// 	fmt.Println(db.Query(sql))
// 	//fmt.Println(pgsql)
// }

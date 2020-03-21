package infra

import (
	"fmt"
	"database/sql"
	"log"
)

const dbname string = "storedb"
const host string = "localhost"
const user string = "postgres"
const password string = "mysecretpassword"

// ConnectToDatabase returns a DB connection and log a fatal error in case of problems
func ConnectToDatabase() *sql.DB {
	log.Println("Connecting to db...")
	connStr := fmt.Sprintf("dbname=%v host=%v port=54320 sslmode=disable user=%v password=%v", dbname, host, user, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

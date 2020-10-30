package database

import (
	"database/sql"

	// Postgres Driver
	_ "github.com/lib/pq"
)

// OpenConnectionDatabase for open connection with database
func OpenConnectionDatabase() *sql.DB {
	credentials := "user=postgres dbname=shorten password=shan host=172.21.0.2 sslmode=disable"
	connection, err := sql.Open("postgres", credentials)
	
	if err != nil {
		panic(err)
	}

	return connection
}
package database

import (
	"database/sql"
	"fmt"
	"os"
	"shortenUrl/reader/config"

	// Postgres database driver
	_ "github.com/lib/pq"
)

// OpenConnectionDatabase for open connection with database
func OpenConnectionDatabase() *sql.DB {
	user := os.Getenv(config.UserKey)
	dbname := os.Getenv(config.DbNameKey)
	passwd := os.Getenv(config.PasswdKey)
	host := os.Getenv(config.HostKey)
	sslMode := os.Getenv(config.SslModeKey)

	credentialPattern := "user=%s dbname=%s password=%s host=%s sslmode=%s"
	credentials := fmt.Sprintf(credentialPattern, user, dbname, passwd, host, sslMode)
	connection, err := sql.Open("postgres", credentials)
	
	if err != nil {
		panic(err)
	}

	return connection
}
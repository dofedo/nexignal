package config

import (
	"database/sql"
	f "fmt"
	"os"
)

var (
	user     = os.Getenv("NEXIGNAL_USER")
	password = os.Getenv("NEXIGNAL_PASS")
	port     = os.Getenv("NEXIGNAL_PORT")
	host     = os.Getenv("NEXIGNAL_HOST")
	db_name  = os.Getenv("NEXIGNAL_DB_NAME")
	ssl_mode = os.Getenv("NEXIGNAL_SSL_MODE")
)

// Connection to DB(postgresql)
func InitDB() *sql.DB {

	db_info := f.Sprintf("user=%s password=%s port=%s host=%s dbname=%s sslmode=%s",
		user, password, port, host, db_name, ssl_mode)

	db, err := sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	}

	return db
}

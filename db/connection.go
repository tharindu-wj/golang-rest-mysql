package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//database connection
func DBConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "asdf1234"
	dbName := "go_rest_mysql"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

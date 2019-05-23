package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//database connection
func DBConn() (conn *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "asdf1234"
	dbName := "go_rest_mysql"

	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return conn
}

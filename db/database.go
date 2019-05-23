package db

import "database/sql"

//database connection
func DBConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "asdf1234"
	dbName := "go_leafycode"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

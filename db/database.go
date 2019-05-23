package db

import "database/sql"

var conn = DBConn()

func FindAll(table string) *sql.Rows {
	selDB, err := conn.Query("SELECT * FROM " + table + " ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}

	return selDB
}

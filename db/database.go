package db

import "database/sql"

var conn = DBConn()

func FindAll(table string) *sql.Rows {
	result, err := conn.Query("SELECT * FROM " + table + " ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}

	return result
}

func FindBy(table string, key string) *sql.Rows {
	result, err := conn.Query("SELECT * FROM " + table + " WHERE p_id=" + key)
	if err != nil {
		panic(err.Error())
	}

	return result
}

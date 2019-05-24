package db

import (
	"bytes"
	"database/sql"
	"fmt"
)

//connection
var connection = DBConn()

//get all rows from a table
func FindAll(table string) *sql.Rows {
	sqlQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at DESC", table)
	result, err := connection.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}

	return result
}

//get single row from a table
func FindBy(table string, key string) *sql.Rows {
	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=%s", table, key)
	result, err := connection.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}

	return result
}

//save record to a table
func Save(table string, item map[string]string) bool {

	var columns, values bytes.Buffer

	i := 1
	itemLength := len(item)

	for k, v := range item {
		if itemLength >= i {
			values.WriteString("'")
		}
		columns.WriteString(k)
		values.WriteString(v)

		if itemLength > i {
			values.WriteString("'")
			columns.WriteString(",")
			values.WriteString(",")
		} else {
			values.WriteString("'")
		}
		i++
	}

	columnString := columns.String()
	valueString := values.String()

	sqlQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", table, columnString, valueString)
	_, err := connection.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	} else {
		return true
	}

}

//delete row from a table
func Remove(table string, key string) bool {
	sqlQuery := fmt.Sprintf("DELETE  FROM %s  WHERE id=%s", table, key)
	_, err := connection.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}

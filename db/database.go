package db

import (
	"bytes"
	"database/sql"
)

var connection = DBConn()

func FindAll(table string) *sql.Rows {
	result, err := connection.Query("SELECT * FROM " + table + " ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}

	return result
}

func FindBy(table string, key string) *sql.Rows {
	result, err := connection.Query("SELECT * FROM " + table + " WHERE p_id=" + key)
	if err != nil {
		panic(err.Error())
	}

	return result
}

func Save(table string, item map[string]string) bool {

	var columns, values bytes.Buffer

	i := 1
	itemLenght := len(item)

	for k, v := range item {
		if (itemLenght >= i) {
			values.WriteString("'")
		}
		columns.WriteString(k)
		values.WriteString(v)

		if (itemLenght > i) {
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

	_, err := connection.Query("INSERT INTO " + table + "(" + columnString + ") VALUES(" + valueString + ")")
	if err != nil {
		panic(err.Error())
	} else {
		return true
	}

}

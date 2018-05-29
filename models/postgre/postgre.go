package postgre

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
)

func Insert(a interface{}) {

	db, err := OpenDB()
	// 插入數據
	defer db.Close()
	CheckErr(err)

	tableName := getTableName(a)

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	columns := getColumnName(tableName)
	var values []string
	for _, column := range columns {
		for j := 0; j < v.NumField(); j++ {
			var name = strings.ToLower(t.Field(j).Name)
			if name == column {
				values = append(values, v.Field(j).String())
				break
			}
		}
	}
	var insertSQL = "INSERT INTO " + tableName + "(" + strings.Join(columns, ",") + ") VALUES('" + strings.Join(values, "','") + "')"
	_, err = db.Exec(insertSQL)
	CheckErr(err)
	// db.QueryRow(insertSQL)

}

func Update(a interface{}) {
	//更新數據
	db, err := OpenDB()
	CheckErr(err)
	tableName := getTableName(a)

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	keyColumns := getKeyColumn(tableName)
	var updateSQL = "update %v set (%v) = ('%v')"

	updateColumns := getUpdateColumn(tableName)
	var values []string
	for _, column := range updateColumns {
		for j := 0; j < v.NumField(); j++ {
			var name = strings.ToLower(t.Field(j).Name)
			if name == column {
				values = append(values, v.Field(j).String())
				break
			}
		}
	}
	if keyColumns != nil {
		updateSQL = updateSQL + " where 1=1"
		for _, column := range keyColumns {
			for j := 0; j < v.NumField(); j++ {
				var name = strings.ToLower(t.Field(j).Name)
				if name == column {
					updateSQL = updateSQL + " and " + name + " = '" + v.Field(j).String() + "'"
					break
				}
			}
		}
	}

	db.QueryRow(fmt.Sprintf(updateSQL, tableName, strings.Join(updateColumns, ","), strings.Join(values, "','")))
	fmt.Print(updateSQL)
	db.Close()

}
func Delete(a interface{}) {
	//删除數據
	db, err := OpenDB()
	CheckErr(err)

	tableName := getTableName(a)
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	var deleteSQL = "delete from %v"

	keyColumns := getKeyColumn(tableName)
	columns := getColumnName(tableName)

	if keyColumns != nil {
		deleteSQL = deleteSQL + " where 1=1"
		for _, column := range keyColumns {
			for j := 0; j < v.NumField(); j++ {
				var name = strings.ToLower(t.Field(j).Name)
				if name == column {
					deleteSQL = deleteSQL + " and " + name + " = '" + v.Field(j).String() + "'"
					break
				}
			}
		}
	} else {
		var values []string
		var keys []string
		for _, column := range columns {
			for j := 0; j < v.NumField(); j++ {
				var name = strings.ToLower(t.Field(j).Name)
				if name == column && v.Field(j).String() != "" {
					keys = append(keys, name)
					values = append(values, v.Field(j).String())
					break
				}
			}
		}
		if values != nil {
			deleteSQL = deleteSQL + " where 1=1 "
			for i := 0; i < len(values); i++ {
				deleteSQL = deleteSQL + " and " + keys[i] + " = '" + values[i] + "'"
			}
		}
	}
	db.QueryRow(fmt.Sprintf(deleteSQL, tableName))
	db.Close()
}

func OpenDB() (*sql.DB, error) {

	return sql.Open("postgres", "user=postgres password=shacom04 dbname=jacklin sslmode=disable")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getColumnName(tableName string) []string {
	db, err := OpenDB()
	CheckErr(err)
	rows, err := db.Query("select column_name from information_schema.columns where table_name=$1 and column_default is null", tableName)
	var columns []string

	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		columns = append(columns, column)
	}
	return columns
}
func getKeyColumn(tableName string) []string {
	db, err := OpenDB()
	CheckErr(err)
	var selectSQL string = "SELECT a.attname FROM   pg_index i JOIN   pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY(i.indkey) WHERE  i.indrelid = $1::regclass AND i.indisprimary"
	rows, err := db.Query(selectSQL, tableName)
	var columns []string

	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		columns = append(columns, column)
	}
	return columns
}
func getUpdateColumn(tableName string) []string {
	db, err := OpenDB()
	CheckErr(err)
	var selectSQL string = `SELECT
		column_name
	FROM
		information_schema. COLUMNS
	WHERE
		table_name = $1
	AND column_default IS NULL
	AND column_name NOT IN (
		SELECT
			a.attname
		FROM
			pg_index i
		JOIN pg_attribute a ON a.attrelid = i.indrelid
		AND a.attnum = ANY (i.indkey)
		WHERE
			i.indrelid = $1 :: regclass
		AND i.indisprimary
	)`

	rows, err := db.Query(selectSQL, tableName)
	var columns []string

	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		columns = append(columns, column)
	}
	return columns
}
func getTableName(a interface{}) string {
	table := strings.Split(reflect.TypeOf(a).String(), ".")

	tableName := strings.ToLower(table[len(table)-1])
	return tableName
}

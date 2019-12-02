package main

import (
	"DumbTree/units"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:li123456@tcp(127.0.0.1:3306)/app?charset=utf8")
	checkErr(err)

	rows, err := db.Query("select  * from ims_test")
	checkErr(err)

	cols, err := rows.Columns()
	checkErr(err)

	value := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for k, _ := range value {
		scans[k] = &value[k]
	}
	var result []map[string]interface{}
	for rows.Next() {
		rows.Scan(scans...)
		row := map[string]interface{}{}
		for k, v := range value {
			key := cols[k]
			row[key] = string(v)
		}
		result = append(result, row)
	}
	db.Close()

	for rk, rv := range result {
		fmt.Printf("这是第%d行", rk)
		fmt.Println()
		units.Foreach(rv, func(k string, v interface{}) {
			fmt.Println(k, v)
		})
	}

}

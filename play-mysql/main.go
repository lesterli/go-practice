package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name        string `json:name`
	LastUpdated string `json:last_updated`
}

func main() {
	fmt.Println("Go play MySQL")
	// 打开数据库连接
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my-database")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected to MySQL database")

	// 表中插入记录 需已创建test表
	//insertRecord(db)

	// 执行数据库表查询
	//getRawBytes(db)
	getResults(db)
}

func insertRecord(db *sql.DB) {
	insert, err := db.Query("INSERT INTO test VALUES('Lester')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("Successfully inserted into test table")
}

func getRawBytes(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error())
	}

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// 切片存储RawBytes
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// 打印每列
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("Successfully listed records from test table")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}

func getResults(db *sql.DB) {
	results, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User

		err = results.Scan(&user.Name, &user.LastUpdated)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}
	fmt.Println("Successfully listed struct from test table")
}

package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

func GetDBconn() *sql.DB {
	dbName := "go_ms_test"
	fmt.Println("conn info :", dbName)
	db, err := sql.Open("mysql", "golang:welcome1@tcp(127.0.0.1:3306)/go_ms_test")
	if err != nil {
		panic(err.Error)
	}
	return db
}

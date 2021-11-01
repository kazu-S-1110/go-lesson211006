package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql") //ファイルがなければ作ってくれる
	defer Db.Close()

	// make table
	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 		name STRING,
	// 		age INT)
	// 	`

	// add data
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)" //?とすると後で入れてくれる

	_, err := Db.Exec(cmd, "hoge", 20)
	if err != nil {
		log.Fatalln(err)
	}
}

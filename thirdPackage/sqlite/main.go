package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql") //ファイルがなければ作ってくれる
	defer Db.Close()

	// make table
	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 		name STRING,
	// 		age INT)
	// 	`

	// add data
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)" //?とすると後で入れられるようになる
	// _, err := Db.Exec(cmd, "fooo", 40)
	// またSQLインジェクション対策にもなるらしい

	// update data
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// _, err := Db.Exec(cmd, 32, "hoge")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// get data specific
	// cmd := "SELECT * FROM persons where age = ?"
	// //QueryRow get 1record
	// row := Db.QueryRow(cmd, 32)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// get dates
	// cmd := "SELECT * FROM persons"
	// // Query は条件に合うものを全て取得
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() { //ちょといつもと違う書き方なので注意
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Println(p)
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// delete data
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "hoge")
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=test_db password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// Create
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)" //sqliteの違うのはここだけ
	_, err := Db.Exec(cmd, "fooo", 20)
	if err != nil {
		log.Fatalln(err)
	}
}

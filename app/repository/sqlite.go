package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

func Conn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./unico.db")

	if err != nil {
		panic(err)
	}

	initial(db)

	return db
}

func initial(db *sql.DB) {
	// Read file
	file, err := ioutil.ReadFile("./dump.sql")
	if err != nil {
			fmt.Println(err.Error())
	}

	// Execute all
	_, err = db.Exec(string(file))
	if err != nil {
			fmt.Println(err.Error())
	} 
}	

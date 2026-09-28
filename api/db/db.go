package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //underscore means import is needed for this package
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5) // at least 5 connection at all time
}

// var DB *sql.DB

// func InitDB() {
//     var err error
//     DB, err = sql.Open("sqlite3", "api.db")

//     if err != nil {
//         panic("Could not connect to database.")
//     }

//     DB.SetMaxOpenConns(10)
//     DB.SetMaxIdleConns(5)

//     createTables()
// }

package util

import (
	"database/sql"
	"log"
)

type Accuracy = string

const (
	Minute     Accuracy = "Minute"
	FiveMinute Accuracy = "FiveMinute"
	Hour       Accuracy = "Hour"
	Day        Accuracy = "Day"
	Week       Accuracy = "Week"
	Month      Accuracy = "Month"
	Quarter    Accuracy = "Quarter"
	Year       Accuracy = "Year"
)

func SqliteQuery(db, command string, args ...interface{}) *sql.Rows {
	DB, err := sql.Open("sqlite3", db)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	rows, err := DB.Query(command, args...)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

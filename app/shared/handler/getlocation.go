package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetLocation(email string) (string) {
	db, err := sql.Open("mysql", "root:@tcp(192.168.200.181:3306)/golang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT location FROM user WHERE email = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	
	var location string
	// Query name
	err = stmtOut.QueryRow(email).Scan(&location) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	
	return location
}

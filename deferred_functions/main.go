package main

import "database/sql"

func main() {
	//Exemplo pratico de deferred function
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close()

	rows, _ := db.Query("some sql query here")
	defer rows.Close()
}

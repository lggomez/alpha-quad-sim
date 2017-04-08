package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetClimate(day int) (error, string) {
	var singleQueryResult string
	err := error(nil)

	connectionName := MustGetenv("CLOUDSQL_CONNECTION_NAME")
	user := MustGetenv("CLOUDSQL_USER")
	password := os.Getenv("CLOUDSQL_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	defer db.Close()

	rows, err := db.Query("SELECT * FROM climates WHERE climate_day=?", day)

	defer rows.Close()

	for rows.Next() {
		var climate string
		if err := rows.Scan(&climate); err != nil {
			break
		} else {
			singleQueryResult = climate
		}
		break
	}

	return err, singleQueryResult
}

func SaveClimate(day int, climate string) (error) {
	err := error(nil)

	connectionName := MustGetenv("CLOUDSQL_CONNECTION_NAME")
	user := MustGetenv("CLOUDSQL_USER")
	password := os.Getenv("CLOUDSQL_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	defer db.Close()

	_, err = db.Exec("INSERT INTO climates (climate,climate_day) VALUES(?,?);", climate, day)

	return err
}

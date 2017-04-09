package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// GetClimate - Retrieve the climate data for a given day
func GetClimate(day int) (error, string) {
	var singleQueryResult string
	err := error(nil)

	connectionName, err := GetEnvironmentVariable("CLOUDSQL_CONNECTION_NAME")
	user, err := GetEnvironmentVariable("CLOUDSQL_USER")
	password, err := GetEnvironmentVariable("CLOUDSQL_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	db.Exec("USE climateregistry")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM climates WHERE climate_day=?", day)
	defer rows.Close()

	rows.Next()
	var climate string
	var climate_day int
	var id int
	err = rows.Scan(&climate, &climate_day, &id)

	if err == nil {
		singleQueryResult = climate
	}

	return err, singleQueryResult
}

// SaveClimate - Save a climate object to the database
func SaveClimate(day int, climate string) (error) {
	err := error(nil)

	connectionName, err := GetEnvironmentVariable("CLOUDSQL_CONNECTION_NAME")
	user, err := GetEnvironmentVariable("CLOUDSQL_USER")
	password, err := GetEnvironmentVariable("CLOUDSQL_PASSWORD")

	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))

	if err != nil {
		return err
	}

	db.Exec("USE climateregistry")
	defer db.Close()

	_, err = db.Exec("INSERT INTO climates (climate, climate_day) VALUES(?,?);", climate, day)
	return err
}

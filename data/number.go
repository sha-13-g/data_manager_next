package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Number struct {
	Id      string
	Number  string
	Site_id string
}

func DeleteNumber(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM number WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetNumberBySite(id string) []Number {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM number WHERE "site_id" = $1`
	rows, err := db.Query(query, id)
	checkIfError(err)

	var number Number
	var numbers []Number

	for rows.Next() {
		rows.Scan(&number.Id, &number.Number, &number.Site_id)
		numbers = append(numbers, number)
	}

	return numbers
}

func GetNumbers() []Number {
	db := InitConn()
	defer db.Close()

	query := "SELECT * FROM number"
	rows, err := db.Query(query)
	checkIfError(err)

	var number Number
	var numbers []Number

	for rows.Next() {
		rows.Scan(&number.Id, &number.Number, &number.Site_id)
		numbers = append(numbers, number)
	}

	return numbers

}

func GetNumber(id string) Number {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM number WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var number Number

	if row.Next() {
		row.Scan(&number)
	}

	return number
}

func (number Number) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO number VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, number.Id, number.Number, number.Site_id)
	checkIfError(err)

	return r
}

func (number Number) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE number SET "number" = $1 WHERE "id" = $2`)
	r, err := db.Exec(query, number.Number, number.Id)
	checkIfError(err)

	return r
}

func DeleteNumberBySite(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM number WHERE "site_id" = $1`
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

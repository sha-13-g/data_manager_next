package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Commercial struct {
	Id      int
	Name    string
	Role    string
	Address string
}

func DeleteCommercial(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM commercial WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetCommercials() []Commercial {
	db := InitConn()
	defer db.Close()

	var (
		commercial  Commercial
		commercials []Commercial
	)

	query := "SELECT * FROM commercial"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&commercial.Id, &commercial.Name, &commercial.Role, &commercial.Address)
		commercials = append(commercials, commercial)
	}

	return commercials
}

func GetCommercial(id int) Commercial {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM commercial  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var commercial Commercial

	if row.Next() {
		row.Scan(&commercial.Id, &commercial.Name, &commercial.Role, &commercial.Address)
	}

	return commercial
}

func (commercial Commercial) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO commercial VALUES (DEFAULT, $1, $2, $3)`)
	r, err := db.Exec(query, commercial.Name, commercial.Role, commercial.Address)
	checkIfError(err)

	return r
}

func (commercial Commercial) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE commercial SET "name" = $2, "role" = $3, "address" = $4 WHERE "id" = $1`)
	r, err := db.Exec(query, commercial.Id, commercial.Name, commercial.Role, commercial.Address)
	checkIfError(err)

	return r
}

func GetCommercialByName(name string) Commercial {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM commercial WHERE "name" = $1`
	row, err := db.Query(query, name)
	checkIfError(err)

	var commercial Commercial

	if row.Next() {
		row.Scan(&commercial.Id, &commercial.Name, &commercial.Role, &commercial.Address)
	}

	return commercial
}

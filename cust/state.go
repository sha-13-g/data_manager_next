package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type State struct {
	Id          int
	Name        string
	Description string
}

func DeleteState(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM state WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func Getstates() []State {
	db := InitConn()
	defer db.Close()

	var (
		state  State
		states []State
	)

	query := "SELECT * FROM state"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&state.Id, &state.Name)
		states = append(states, state)
	}

	return states
}

func Getstate(id int) State {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM state  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var state State

	if row.Next() {
		row.Scan(&state.Id, &state.Name)
	}

	return state
}

func (state State) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, &state.Id, &state.Name, &state.Description)
	checkIfError(err)

	return r
}

func (state State) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE state SET "name" = $2, "description" = $2 WHERE "id" = $1`)
	r, err := db.Exec(query, &state.Id, &state.Name, &state.Description)
	checkIfError(err)

	return r
}

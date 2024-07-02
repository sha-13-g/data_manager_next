package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Step struct {
	Id          int
	Name        string
	Description string
}

func DeleteStep(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM step WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetSteps() []Step {
	db := InitConn()
	defer db.Close()

	var (
		step  Step
		steps []Step
	)

	query := "SELECT * FROM step"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&step.Id, &step.Name)
		steps = append(steps, step)
	}

	return steps
}

func GetStep(id int) Step {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM step  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var step Step

	if row.Next() {
		row.Scan(&step.Id, &step.Name)
	}

	return step
}

func (step Step) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, &step.Id, &step.Name, &step.Description)
	checkIfError(err)

	return r
}

func (step Step) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE step SET "name" = $2, "description" = $2 WHERE "id" = $1`)
	r, err := db.Exec(query, &step.Id, &step.Name, &step.Description)
	checkIfError(err)

	return r
}

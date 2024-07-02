package cust

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type ProjectState struct {
	Project_id int
	State_id   int
	Date       time.Time
}

func DeleteProjectState(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM project_state WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func Getproject_states() []ProjectState {
	db := InitConn()
	defer db.Close()

	var (
		project_state  ProjectState
		project_states []ProjectState
	)

	query := "SELECT * FROM project_state"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&project_state.Project_id, &project_state.State_id, &project_state.Date)
		project_states = append(project_states, project_state)
	}

	return project_states
}

func GetProjectState(id int) ProjectState {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM project_state  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var project_state ProjectState

	if row.Next() {
		row.Scan(&project_state.Project_id, &project_state.State_id, &project_state.Date)
	}

	return project_state
}

func (project_state ProjectState) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO  VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, &project_state.Project_id, &project_state.State_id, &project_state.Date)
	checkIfError(err)

	return r
}

func (project_state ProjectState) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE project_state SET "project_id" = $2, "state_id" = $3,  WHERE "id" = $1`)
	r, err := db.Exec(query, &project_state.Project_id, &project_state.State_id, &project_state.Date)
	checkIfError(err)

	return r
}

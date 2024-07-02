package data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type ProjectStep struct {
	ProjectId int
	StepId    int
	Date      time.Time
}

func DeleteProjectStep(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM project_step WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetprojectSteps() []ProjectStep {
	db := InitConn()
	defer db.Close()

	var (
		projectStep  ProjectStep
		projectSteps []ProjectStep
	)

	query := "SELECT * FROM project_step"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&projectStep.ProjectId, &projectStep.StepId, &projectStep.Date)
		projectSteps = append(projectSteps, projectStep)
	}

	return projectSteps
}

func GetProjectStep(id int) ProjectStep {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM project_step  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var projectStep ProjectStep

	if row.Next() {
		row.Scan(&projectStep.ProjectId, &projectStep.StepId, &projectStep.Date)
	}

	return projectStep
}

func (projectStep ProjectStep) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO  VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, &projectStep.ProjectId, &projectStep.StepId, &projectStep.Date)
	checkIfError(err)

	return r
}

func (projectStep ProjectStep) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE project_step SET "project_id" = $2, "step_id" = $3,  WHERE "id" = $1`)
	r, err := db.Exec(query, &projectStep.ProjectId, &projectStep.StepId, &projectStep.Date)
	checkIfError(err)

	return r
}

package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Project struct {
	Id            int
	Name          string
	Description   string
	Customer_id   string
	Commercial_id string
	Service_id    string
}

type ProjectDetails struct {
	Project_id       int
	Project_name     string
	Customer_name    string
	Customer_branch  string
	Customer_contact string
	Service_name     string
	Commercial_last  string
	Commercial_first string
}

func DeleteProject(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM project WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetProjects() []Project {
	db := InitConn()
	defer db.Close()

	var (
		project  Project
		projects []Project
	)

	query := "SELECT * FROM project"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&project.Id, &project.Name, &project.Description, &project.Customer_id, &project.Commercial_id, &project.Service_id)
		projects = append(projects, project)
	}

	return projects
}

func GetProject(id int) Project {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM project  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var project Project

	if row.Next() {
		row.Scan(&project.Id, &project.Name, &project.Description, &project.Customer_id, &project.Commercial_id, &project.Service_id)
	}

	return project
}

func (project Project) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO project VALUES (DEFAULT, $1, $2, $3, $4, $5)`)
	r, err := db.Exec(query, &project.Name, &project.Description, &project.Customer_id, &project.Commercial_id, &project.Service_id)
	checkIfError(err)

	return r
}

func (project Project) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE project SET "name" = $2, "description" = $3, "customer_id" = $4, "commercial_id" = $5, "service_id" = $6 WHERE "id" = $1`)
	r, err := db.Exec(query, &project.Id, &project.Name, &project.Description, &project.Customer_id, &project.Commercial_id, &project.Service_id)
	checkIfError(err)

	return r
}

func GetProjectByAll() []ProjectDetails {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`SELECT p.id, p.name, cs.name cs_name, cs.branch, cs.contact, s.name s_name, c.name
	FROM project p, customer cs, service s, commercial c 
	WHERE p.customer_id = cs.id AND p.service_id = s.id AND cs.commercial_id = c.id`)

	r, err := db.Query(query)
	checkIfError(err)

	var projects []ProjectDetails
	var project ProjectDetails

	for r.Next() {
		r.Scan(&project.Project_id, &project.Project_name, &project.Customer_name, &project.Customer_branch,
			&project.Customer_contact, &project.Service_name, &project.Commercial_last, &project.Commercial_first)
		projects = append(projects, project)
	}

	return projects
}

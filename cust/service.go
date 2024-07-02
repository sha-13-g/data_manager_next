package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Service struct {
	Id   int
	Name string
}

func DeleteService(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM service WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetServices() []Service {
	db := InitConn()
	defer db.Close()

	var (
		service  Service
		services []Service
	)

	query := "SELECT * FROM service"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&service.Id, &service.Name)
		services = append(services, service)
	}

	return services
}

func GetService(id int) Service {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM service  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var service Service

	if row.Next() {
		row.Scan(&service.Id, &service.Name)
	}

	return service
}

func (service Service) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO service VALUES (DEFAULT, $1)`)
	r, err := db.Exec(query, &service.Name)
	checkIfError(err)

	return r
}

func (service Service) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE service SET "name" = $2 WHERE "id" = $1`)
	r, err := db.Exec(query, &service.Id, &service.Name)
	checkIfError(err)

	return r
}

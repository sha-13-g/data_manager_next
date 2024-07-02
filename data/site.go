package data

import (
	"database/sql"
	"fmt"
)

type Site struct {
	Id          string
	Name        string
	Customer_id string
}

func DeleteSite(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM site WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetSitesByCustomer(id string) []Site {
	var sites []Site
	var site Site
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM site WHERE "customer_id" = $1`
	rows, err := db.Query(query, id)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&site.Id, &site.Name, &site.Customer_id)
		sites = append(sites, site)
	}

	return sites
}

func GetSites() []Site {
	var sites []Site
	var site Site
	db := InitConn()
	defer db.Close()

	query := "SELECT * FROM site"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&site.Id, &site.Name, &site.Customer_id)
		sites = append(sites, site)
	}

	return sites
}

func GetSite(id string) Site {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM Site WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var site Site

	if row.Next() {
		row.Scan(&site)
	}

	return site
}

func (site Site) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO site VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, site.Id, site.Name, site.Customer_id)
	checkIfError(err)

	return r
}

func (site Site) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE site SET "name" = $1 WHERE "id" = $2`)
	r, err := db.Exec(query, site.Name, site.Id)
	checkIfError(err)

	return r
}

func DeleteSitesByCustomer(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM site WHERE "customer_id" = $1`
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

func GetSiteByName(name string) Site {
	var site Site
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM site WHERE "name" = $1`
	row, err := db.Query(query, name)
	checkIfError(err)

	for row.Next() {
		row.Scan(&site.Id, &site.Name, &site.Customer_id)
	}

	return site
}

package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Prospect struct {
	Id               int
	Name             string
	Address          string
	Referant         string
	Referant_contact string
	Contact          string
	Mail_object      string
	Ddp              string
	Ddc              string
}

func DeleteProspect(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM prospect WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetProspects() []Prospect {
	db := InitConn()
	defer db.Close()

	var (
		prospect  Prospect
		prospects []Prospect
	)

	query := "SELECT * FROM prospect"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&prospect.Id, &prospect.Name, &prospect.Address, &prospect.Referant, &prospect.Referant_contact, &prospect.Contact, &prospect.Mail_object, &prospect.Ddp, &prospect.Ddc)

		prospect.Ddp = prospect.Ddp[:10]
		prospect.Ddc = prospect.Ddc[:10]

		prospects = append(prospects, prospect)
	}

	return prospects
}

func Getprospect(id int) Prospect {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM prospect  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var prospect Prospect

	if row.Next() {
		row.Scan(&prospect.Id, &prospect.Name, &prospect.Address, &prospect.Referant, &prospect.Referant_contact, &prospect.Contact, &prospect.Mail_object, &prospect.Ddp, &prospect.Ddc)
	}

	return prospect
}

func (prospect Prospect) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO prospect VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)`)
	r, err := db.Exec(query, prospect.Name, prospect.Address, prospect.Referant, prospect.Referant_contact, prospect.Contact, prospect.Mail_object, prospect.Ddp, prospect.Ddc)
	checkIfError(err)

	return r
}

func (prospect Prospect) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE prospect SET "name" = $2, "address" = $3, "referant" = $4, "referant_contact" = $5, "contact" = $6, "mail_object" = $7, "ddp" = $8, "ddc" = $9 WHERE "id" = $1`)
	r, err := db.Exec(query, prospect.Id, prospect.Name, prospect.Address, prospect.Referant, prospect.Referant_contact, prospect.Contact, prospect.Mail_object, prospect.Ddp, prospect.Ddc)
	checkIfError(err)

	return r
}

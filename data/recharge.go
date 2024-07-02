package data

import (
	"database/sql"
	"fmt"
	"time"
)

type Recharge struct {
	Id        string
	Volume    int
	Date_re   time.Time
	Date_exp  time.Time
	Auto_re   bool
	Number_id string
	Volume_id int
}

func (recharge Recharge) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO recharge VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	r, err := db.Exec(query, recharge.Id, recharge.Volume, recharge.Date_re, recharge.Date_exp, recharge.Auto_re, recharge.Number_id, recharge.Volume_id)
	checkIfError(err)

	return r
}

func (recharge Recharge) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE recharge SET "volume" = $2, "date_re" = $3, "date_exp" = $4 WHERE "id" = $1`)
	r, err := db.Exec(query, recharge.Id, recharge.Volume, recharge.Date_re, recharge.Date_exp)
	checkIfError(err)

	return r
}

func DeleteRecharge(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM recharge WHERE "id" = $1`)
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

func GetRecharge(id string) Recharge {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM recharge WHERE "id" = $1`)
	row, err := db.Query(query, id)
	checkIfError(err)

	var recharge Recharge
	if row.Next() {
		row.Scan(&recharge.Id, &recharge.Volume, &recharge.Date_re, &recharge.Date_exp, &recharge.Auto_re, &recharge.Number_id, &recharge.Volume_id)
	}

	return recharge
}

func GetRecharges() []Recharge {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM recharge")
	rows, err := db.Query(query)
	checkIfError(err)

	var recharge Recharge
	var recharges []Recharge

	for rows.Next() {
		rows.Scan(&recharge.Id, &recharge.Volume, &recharge.Date_re, &recharge.Date_exp, &recharge.Auto_re, &recharge.Number_id, &recharge.Volume_id)
		recharges = append(recharges, recharge)
	}

	return recharges
}

func GetRechargesByNumber(id string) []Recharge {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM recharge WHERE "number_id" = $1`)
	rows, err := db.Query(query, id)
	checkIfError(err)

	var recharge Recharge
	var recharges []Recharge

	for rows.Next() {
		rows.Scan(&recharge.Id, &recharge.Volume, &recharge.Date_re, &recharge.Date_exp, &recharge.Auto_re, &recharge.Number_id, &recharge.Volume_id)
		recharges = append(recharges, recharge)
	}

	return recharges
}

func DeleteRechargesByNumber(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM recharge WHERE "number_id" = $1`)
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

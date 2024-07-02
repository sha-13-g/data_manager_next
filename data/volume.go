package data

import (
	"database/sql"
	"fmt"
)

type Volume struct {
	Id     int
	Volume int
	Price  string
}

func (volume Volume) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO volume VALUES ($1, $2, $3)`)
	r, err := db.Exec(query, volume.Id, volume.Volume, volume.Price)
	checkIfError(err)

	return r
}

func (volume Volume) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE volume SET "volume" = $2, "price" = $3 WHERE "id" = $1`)
	r, err := db.Exec(query, volume.Id, volume.Volume, volume.Price)
	checkIfError(err)

	return r
}

func DeleteVolume(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM volume WHERE "id" = $1`)
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

func GetVolume(id int) Volume {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM volume WHERE "id" = $1`)
	row, err := db.Query(query, id)
	checkIfError(err)

	var volume Volume
	if row.Next() {
		row.Scan(&volume.Id, &volume.Volume, &volume.Price)
	}

	return volume
}

func GetVolumes() []Volume {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM volume")
	rows, err := db.Query(query)
	checkIfError(err)

	var volume Volume
	var volumes []Volume

	for rows.Next() {
		rows.Scan(&volume.Id, &volume.Volume, &volume.Price)
		volumes = append(volumes, volume)
	}

	return volumes
}

package data

import (
	"database/sql"
	"fmt"
	"time"
)

type Incident struct {
	Id             string
	Description    string
	Origin         int
	Responsability string
	Date           time.Time
	CustomerId     int
	AgentId        int
	Site_id        string
}

type IncidentTable struct {
	Id             string
	Description    string
	Origin         string
	Site           string
	Responsability string
	Name           string
	Date           string
}

func (incident Incident) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO incident VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7)`)
	r, err := db.Exec(query, incident.Description, incident.Origin, incident.Responsability, incident.Date, incident.CustomerId, incident.AgentId, incident.Site_id)

	checkIfError(err)

	return r
}

func (incident Incident) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE incident SET "description" = $2, "origin" = $3, "responsability" = $4, "date" = $5, 
	"customer_id" = $6, "agent_id" = $7, "site_id" = $8 WHERE "id" = $1 `)

	r, err := db.Exec(query, incident.Id, incident.Description, incident.Origin, incident.Responsability, incident.Date, incident.CustomerId, incident.AgentId, incident.Site_id)

	checkIfError(err)

	return r
}

func GetIncidentTable() []IncidentTable {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`select i.id, i.description, c.name, s.name, i.responsability, a.name, i.date 
	from incident i, customer c, commercial a, site s 
	where c.id = i.customer_id and a.id = i.agent_id and i.site_id = s.id`)

	rows, err := db.Query(query)

	checkIfError(err)

	var incidents []IncidentTable
	var incident IncidentTable

	for rows.Next() {
		rows.Scan(&incident.Id, &incident.Description, &incident.Origin, &incident.Site, &incident.Responsability, &incident.Name, &incident.Date)

		incident.Date = incident.Date[:10]

		incidents = append(incidents, incident)
	}

	return incidents
}

func DeleteIncident(id string) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM incident WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

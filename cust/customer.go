package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id            int
	Logo          string
	Name          string
	Address       string
	Contact       string
	Branch        string
	Commercial_id string
}

func DeleteCustomer(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM customer WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetCustomers() []Customer {
	db := InitConn()
	defer db.Close()

	var (
		customer  Customer
		customers []Customer
	)

	query := `SELECT * FROM customer`
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&customer.Id, &customer.Logo, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
		customers = append(customers, customer)
	}

	return customers
}

func GetCustomer(id int) Customer {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM customer  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var customer Customer

	if row.Next() {
		row.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
	}

	return customer
}

func (customer Customer) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO customer VALUES (DEFAULT, $1, $2, $3, $4, $5, $6)`)
	r, err := db.Exec(query, &customer.Logo, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
	checkIfError(err)

	return r
}

func (customer Customer) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE customer SET "name" = $2, "address" = $3, "contact" = $4, "branch" = $5 "commercial_id" = $6 WHERE "id" = $1`)
	r, err := db.Exec(query, &customer.Id, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
	checkIfError(err)

	return r
}

package data

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
	Commercial_id int
}

type table struct {
	Customer_id   string
	Customer_name string
	Site_id       string
	Site_name     string
	Number_id     string
	Number_name   string
	Recharge_id   string
	Volume        int
	Date_re       string
	Date_exp      string
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

	query := fmt.Sprintf(`UPDATE customer SET "name" = $2, "address" = $3, "contact" = $4, "branch" = $5, "commercial_id" = $6 WHERE "id" = $1`)
	r, err := db.Exec(query, &customer.Id, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
	checkIfError(err)

	return r
}

func DeleteCustomer(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM customer customer WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetCustomersDetails() []table {

	var row table
	var tables []table

	db := InitConn()
	defer db.Close()

	query := fmt.Sprint(`select c.id, c.name, s.id, s.name, n.id, n.number, r.id, r.volume, r.date_re, r.date_exp 
	from customer c, site s, number n, recharge r
	where c.id = s.customer_id and s.id = n.site_id and n.id = r.number_id`)

	r, err := db.Query(query)
	checkIfError(err)

	for r.Next() {

		r.Scan(&row.Customer_id, &row.Customer_name, &row.Site_id, &row.Site_name, &row.Number_id, &row.Number_name, &row.Recharge_id, &row.Volume, &row.Date_re, &row.Date_exp)

		//row.Date_re, _ = time.Parse("2006-01-02", row.Date_re.Format("2006-01-02"))
		//row.Date_exp, _ = time.Parse("2006-01-02", row.Date_exp.Format("2006-01-02"))

		//fmt.Println(time.Parse("2006-01-02", row.Date_exp.Format("2006-01-02")))

		row.Date_re = row.Date_re[:10]
		row.Date_exp = row.Date_exp[:10]

		tables = append(tables, row)

	}

	return tables
}

func GetCustomerByName(name string) Customer {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM customer  WHERE "name" = $1`
	row, err := db.Query(query, name)
	checkIfError(err)

	var customer Customer

	if row.Next() {
		row.Scan(&customer.Id, &customer.Logo, &customer.Name, &customer.Address, &customer.Contact, &customer.Branch, &customer.Commercial_id)
	}

	return customer
}

func UpdateCommercialId(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `UPDATE customer SET "commercial_id" = $1`
	r, err := db.Exec(query, id)
	checkIfError(err)

	return r
}

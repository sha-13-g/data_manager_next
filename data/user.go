package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id            int
	Email         string
	Password      string
	Commercial_id int
}

func DeleteUser(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM users WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetUsers() []User {
	db := InitConn()
	defer db.Close()

	var (
		user  User
		users []User
	)

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Password, &user.Commercial_id)
		users = append(users, user)
	}

	return users
}

func GetUser(id int) User {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM users  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var user User

	if row.Next() {
		row.Scan(&user.Id, &user.Email, &user.Password)
	}

	return user
}

func (user User) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO users VALUES (DEFAULT, $1, $2, $3)`)
	r, err := db.Exec(query, &user.Email, &user.Password, &user.Commercial_id)
	checkIfError(err)

	return r
}

func (user User) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE users SET "email" = $2, "password" = $3, "commercial_id" = $4 WHERE "id" = $1`)
	r, err := db.Exec(query, &user.Id, &user.Email, &user.Password, &user.Commercial_id)
	checkIfError(err)

	return r
}

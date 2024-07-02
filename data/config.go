package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gbl"
	password = "giaitbl"
	dbname   = "data_manager"
)

func checkIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitConn() *sql.DB {
	var psqlconn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	checkIfError(err)

	err = db.Ping()
	checkIfError(err)

	return db
}

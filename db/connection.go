package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {

	dbSource := GetConnectionString()

	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}

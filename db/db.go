package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaDb() *sql.DB {
	// conexao := "postgres://postgres:123456789@localhost/alura_loja?sslmode=disable"
	conexao := "user=postgres dbname=alura_loja password=123456789 host=localhost port=54499 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

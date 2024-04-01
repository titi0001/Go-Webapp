package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaDb() (*sql.DB, error) {
	// conexao := "postgres://postgres:123456789@localhost/alura_loja?sslmode=disable"
	conexao := "user=postgres dbname=alura_loja password=123456789 host=localhost port=54499 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
		return nil, err
	}
	log.Println("Conexão bem-sucedida com o banco de dados!")
	return db, nil
}

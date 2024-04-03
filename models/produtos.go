package models

import (
	"github.com/titi0001/Go-REST-ORM/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() ([]Produto, error) {
	db, err := db.ConectaDb()
	if err != nil {
		return nil, err // Retornar erro de consulta
	}

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		return nil, err // Retornar erro de consulta
	}
	defer selectDeTodosOsProdutos.Close()

	var produtos []Produto

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			return nil, err // Retornar erro de leitura
		}

		p := Produto{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos, nil
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int)  {
	db, _ := db.ConectaDb()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

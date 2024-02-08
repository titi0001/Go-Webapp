package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"log"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}


func conectaDb() (*sql.DB, error) {
	// conexao := "postgres://postgres:123456789@localhost/alura_loja?sslmode=disable"
	conexao := "user=postgres dbname=alura_loja password=123456789 host=localhost port=42381 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
		return nil, err
	}
	log.Println("Conexão bem-sucedida com o banco de dados!")
	return db, nil
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db, err := conectaDb()
	if err != nil {
		// Tratar erro de conexão
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	selectDeTodosProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		// Tratar erro de consulta
		http.Error(w, "Erro ao consultar o banco de dados", http.StatusInternalServerError)
		return
	}
	defer selectDeTodosProdutos.Close()

	var produtos []Produto

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			// Tratar erro de leitura
			http.Error(w, "Erro ao ler os dados do banco de dados", http.StatusInternalServerError)
			return
		}

		p := Produto{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}

		produtos = append(produtos, p)
	}

	err = temp.ExecuteTemplate(w, "index", produtos)
	if err != nil {
		// Tratar erro de execução do template
		http.Error(w, "Erro ao executar o template", http.StatusInternalServerError)
		return
	}
}


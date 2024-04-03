package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/titi0001/Go-REST-ORM/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos, err := models.BuscaTodosOsProdutos()
	if err != nil {
		// Tratar erro de busca de produtos
		http.Error(w, "Erro ao buscar os produtos", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, float64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeConvertido, err := strconv.Atoi(quantidade, int)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", 301)

}

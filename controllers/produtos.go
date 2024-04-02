package controllers

import (
	"net/http"
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

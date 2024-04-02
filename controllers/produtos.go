package controllers

import (
	"net/http"
	"text/template"

	"github.com/titi0001/Go-REST-ORM/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

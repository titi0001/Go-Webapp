package routes

import (
	"net/http"
	"github.com/titi0001/Go-REST-ORM/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}

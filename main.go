package main

import (
	"net/http"
	"github.com/titi0001/Go-REST-ORM/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":9000", nil)
}

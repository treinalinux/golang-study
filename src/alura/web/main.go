package main

import (
	"alura/web/routes"
	"net/http"
)

// main com a chamada para o servidor http
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

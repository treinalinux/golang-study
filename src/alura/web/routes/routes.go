package routes

import (
	"alura/web/controllers"
	"net/http"
)

// CarregaRotas do web
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}

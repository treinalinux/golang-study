package routes

import (
	"alura/web/controllers"
	"net/http"
)

// CarregaRotas do web
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}

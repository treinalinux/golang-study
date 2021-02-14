package main

import (
	"net/http"
	"text/template"
)

// Produto usando uma estrutura para a loja Alura
type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 25, Quantidade: 10},
		{Nome: "Camiseta", Descricao: "Verde, bem bonita", Preco: 55, Quantidade: 50},
		{Nome: "Camiseta", Descricao: "Preta, bem bonita", Preco: 45, Quantidade: 70},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

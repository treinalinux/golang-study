package controllers

import (
	"alura/web/models"
	"html/template"
	"net/http"
)

// temp para chamar meus templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index para chamar minha struct
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

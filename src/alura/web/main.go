package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=alan dbname=alura_loja password=12345678 host=devops.empresa.corp sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Produto usando uma estrutura para a loja Alura
type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

// temp para chamar meus templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

// main com a chamada para o servidor http
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index para chamar minha struct
func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select *from produtos")
	if err != nil {
		panic(err.Error())
	}

	// p será a instancia de 1 produto
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

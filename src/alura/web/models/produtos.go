package models

import (
	"alura/web/db"
)

// Produto usando uma estrutura para a loja Alura
type Produto struct {
	ID              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

// BuscaTodosOsProdutos no banco de dados
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

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

	defer db.Close()
	return produtos
}

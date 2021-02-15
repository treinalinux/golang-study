package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados postgres
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=alan dbname=alura_loja password=12345678 host=devops.empresa.corp sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func conectaComBanco() *sql.DB {
	connStr := "host=localhost user=postgre password=12345678 dbname=alura_loja port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("falha ao conectar com o banco")
		log.Fatal(err)
	}
	fmt.Println("conenctado com banco")
	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBanco()

	rows, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}
	fmt.Println("Chegou no for")
	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
		produtos = append(produtos, produto)

	}

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
	fmt.Println("fechou conexao com o banco")
}

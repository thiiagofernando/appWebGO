package models

import (
	"appwebgo/db"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	conect := db.ConectaComBanco()

	rows, err := conect.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}
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
	defer conect.Close()
	return produtos
}

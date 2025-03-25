package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	connStr := "host=localhost user=postgres password=1234567 dbname=alura_loja port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("falha ao conectar com o banco")
		log.Fatal(err)
	}
	fmt.Println("conenctado com banco")
	return db
}

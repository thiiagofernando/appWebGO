package controllers

import (
	"appwebgo/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()

	temp.ExecuteTemplate(w, "index", produtos)
}

func Novo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "novo", nil)
}

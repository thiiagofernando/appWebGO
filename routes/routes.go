package routes

import (
	"appwebgo/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novo", controllers.Novo)
}

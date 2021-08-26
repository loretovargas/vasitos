package main

import (
	"net/http"
	"ruleta/adapter/in/web"
	"ruleta/service"
)

func main() {

	girarRuletaService := service.GirarRuletaService{}

	girarRuletaHandler := web.GirarRuletaHandler{
		GirarRuletaPortIN: girarRuletaService.GirarRuleta,
	}

	http.HandleFunc("/girarRuleta", girarRuletaHandler.GirarRuleta)

	http.ListenAndServe(":8081", nil)
}

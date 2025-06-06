package main

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/src/utils"
	"fmt"

	"log"
	"net/http"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}

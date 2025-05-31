package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasHome = []Rota{
	{
		URI:                "/home",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},
}

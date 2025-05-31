package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasLogin = []Rota{
	{
		URI:                "/",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.FazerLogin,
		RequerAutenticacao: false,
	},
}

package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotaLogout = Rota{

	URI:                "/logout",
	MetodoHttp:         http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}

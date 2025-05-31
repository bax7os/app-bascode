package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuarios",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscar-usuarios",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeBuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/parar-de-seguir",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/perfil",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuarioLogado,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-usuario",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-usuario",
		MetodoHttp:         http.MethodPut,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizar-senha",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeAtualizacaoDeSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizar-senha",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletar-usuario",
		MetodoHttp:         http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
}

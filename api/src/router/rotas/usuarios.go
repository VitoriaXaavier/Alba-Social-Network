package rotas

import (
	"VitoriaXaavier/Alba-Social-Network/src/controllers"
	"net/http"
)

var rotasUsuarios  = []Rota{
	//Cria um usuario
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	//Busca todos os usuarios
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscaTodosUsuarios,
		RequerAutenticacao: false,
	},
	//Buscando usuario por ID
	{
		URI: "/usuario/{usuarioID}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscandoUsuarioPorID,
		RequerAutenticacao: false,
	},
	// Atualiza usuario
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizaUsuario,
		RequerAutenticacao: false,
	},
	// Deleta usuario
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletaUsuario,
		RequerAutenticacao: false,
	},

}
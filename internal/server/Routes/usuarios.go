package router

import (
	controllers "api/internal/controllers"
	"net/http"
)

var rotaUsuarios = []Rota{
	// Create
	{
		URI:         "/usuarios",
		Metodo:      http.MethodPost,
		Funcao:      controllers.CriarUsuario,
		RequireAuth: false,
	},

	// Request all
	{
		URI:         "/usuarios",
		Metodo:      http.MethodGet,
		Funcao:      controllers.BuscarUsuarios,
		RequireAuth: false,
	},

	// Request one
	{
		URI:         "/usuario/{userId}",
		Metodo:      http.MethodGet,
		Funcao:      controllers.BuscarUsuario,
		RequireAuth: false,
	},

	// UpDate
	{
		URI:         "/usuario/{usuarioId}",
		Metodo:      http.MethodPut,
		Funcao:      controllers.AtualizarUsuario,
		RequireAuth: false,
	},

	// Delete
	{
		URI:         "/usuario/{usuarioId}",
		Metodo:      http.MethodDelete,
		Funcao:      controllers.DeleteUsuario,
		RequireAuth: false,
	},
}

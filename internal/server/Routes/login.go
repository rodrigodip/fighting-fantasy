package router

import (
	controllers "api/internal/controllers"
	"net/http"
)

var loginRoute = Rota{
	URI:         "/login",
	Metodo:      http.MethodPost,
	Funcao:      controllers.Login,
	RequireAuth: false,
}

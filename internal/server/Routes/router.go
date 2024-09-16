package router

import (
	"api/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas as rotas do Sistema
type Rota struct {
	URI         string
	Metodo      string
	Funcao      func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Config(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios
	rotas = append(rotas, loginRoute)

	for _, rota := range rotas {

		if rota.RequireAuth {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Authentication(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}

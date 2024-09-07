package router

import (
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

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

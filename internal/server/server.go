package server

import (
	router "api/internal/server/Routes"

	"github.com/gorilla/mux"
)

func Up() *mux.Router {
	r := mux.NewRouter()
	return router.Config(r)
}

package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/auth/").Handler(http.StripPrefix("/auth", AuthRoutes()))

	return router
}

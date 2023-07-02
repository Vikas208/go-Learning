package routers

import (
	"github.com/Vikas208/social-media-api/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes() *mux.Router {
	AuthRouter := mux.NewRouter()
	AuthRouter.HandleFunc("/signup", controllers.Signup).Methods("POST")
	return AuthRouter
}

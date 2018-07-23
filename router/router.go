package router

import (
	"github.com/go-squads/genrevan-scheduler/controller"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/lxc", controller.GetLXCs).Methods("GET")
	router.HandleFunc("/lxc/{id}", controller.GetLXC).Methods("GET")
	router.HandleFunc("/lxc", controller.CreateLXC).Methods("POST")
	router.HandleFunc("/lxc/{id}/state", controller.UpdateLXCState).Methods("PATCH")

	return router
}

package router

import (
	"net/http"
	
	"github.com/go-squads/genrevan-scheduler/controller"
	"github.com/gorilla/mux"
)

func withAllowCORS(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Max-Age", "86400")
				controller.RespondWithJSON(w, http.StatusOK, nil)
			}
			next.ServeHTTP(w,r)
		}
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/lxc", controller.GetLXCs).Methods("GET")
	router.HandleFunc("/lxc/{id}", controller.GetLXC).Methods("GET")
	router.HandleFunc("/lxc", controller.CreateLXC).Methods("POST")
	router.HandleFunc("/lxc/{id}/state", withAllowCORS(controller.UpdateLXCState)).Methods("PATCH")
	router.HandleFunc("/lxc/{id}/ip", controller.UpdateLXCIp).Methods("PATCH")
	router.HandleFunc("/lxc/{id}", controller.DeleteLXC).Methods("DELETE")
	router.HandleFunc("/lxc/lxd/{id}", controller.GetLXCsByLXDId).Methods("GET")

	router.HandleFunc("/lxd", controller.GetLXDs).Methods("GET")
	router.HandleFunc("/lxd/register", controller.RegisterLXD).Methods("GET")
	router.HandleFunc("/lxd/{id}", controller.GetLXD).Methods("GET")

	router.HandleFunc("/metric/{id}", controller.UpdateMetric).Methods("PUT")

	return router
}

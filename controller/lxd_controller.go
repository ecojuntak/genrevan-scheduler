package controller

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/gorilla/mux"
)

var lxdModel model.Lxd

func GetLXDs(w http.ResponseWriter, r *http.Request) {
	lxds, err := lxdModel.GetLXDs()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxds)
}

func GetLXD(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ip := params["ip"]

	lxd, err := lxdModel.GetLXD(ip)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxd)
}

func RegisterLXD(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Printf("userip: %q is not IP:port", r.RemoteAddr)
	}

	if len(ip) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Ip cannot be empty")
	}

	id, err := lxdModel.CreateLXD(ip)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	data := make(map[string]int)
	data["id"] = *id

	RespondWithJSON(w, http.StatusCreated, data)
}

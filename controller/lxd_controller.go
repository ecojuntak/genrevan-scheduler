package controller

import (
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
	ip := r.FormValue("ip")

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

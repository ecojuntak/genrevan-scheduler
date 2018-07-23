package controller

import (
	"net/http"
	"strconv"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/gorilla/mux"
)

var lxcModel model.Lxc

func GetLXCs(w http.ResponseWriter, r *http.Request) {
	lxcs, err := lxcModel.GetLXCs()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxcs)
}

func GetLXC(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	lxc, err := lxcModel.GetLXC(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxc)
}

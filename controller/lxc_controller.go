package controller

import (
	"net/http"

	"github.com/go-squads/genrevan-scheduler/model"
)

var lxcModel model.Lxc

func GetLXCs(w http.ResponseWriter, r *http.Request) {
	lxcs, err := lxcModel.GetLXCs()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxcs)
}

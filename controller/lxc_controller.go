package controller

import (
	"net/http"
	"strconv"

	"github.com/guregu/null"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/scheduler"
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

func CreateLXC(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	image := r.FormValue("image")

	if len(name) == 0 || len(image) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Name or Image cannot be empty")
	}

	metric, err := scheduler.GetLowestMetricLoad()

	lxc := model.Lxc{
		Name:  name,
		Image: image,
		LxdId: null.NewInt(int64(metric.IdLxd), true),
	}

	id, err := lxcModel.CreateLXC(lxc)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusCreated, id)
}

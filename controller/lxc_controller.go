package controller

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/guregu/null"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/scheduler"
	"github.com/gorilla/mux"
)

var lxcModel model.Lxc

const (
	minimumPortNum = 2000
	maximumPortNum = 3000
	maxIterationsFindHostPort = 15
)

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
	containerPortStr := r.FormValue("containerPort")

	if !isLXCFormValid(name, image, containerPortStr) {
		RespondWithError(w, http.StatusBadRequest, "Name, Image, or Container Port cannot be empty")
	}

	containerPort, err := strconv.Atoi(containerPortStr)

	metric, err := scheduler.GetLowestMetricLoad()

	hostPort, err := getValidHostPort(metric.IdLxd)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		lxc := model.Lxc{
			Name:          name,
			Image:         image,
			LxdId:         null.NewInt(int64(metric.IdLxd), true),
			HostPort:      hostPort,
			ContainerPort: containerPort,
		}
	
		id, err := lxcModel.CreateLXC(lxc)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			RespondWithJSON(w, http.StatusCreated, id)
		}
	}	
}

func getValidHostPort(idLXD int) (int, error) {
	var (
		hostPort int
		isExist bool = true
		count int = 0
		err error
	)
	for isExist && count < maxIterationsFindHostPort {
		hostPort = getRandomPortNumber(minimumPortNum, maximumPortNum)
		isExist, err = lxcModel.IsLXCsExist(idLXD, hostPort)
		if err != nil {
			return 0, err
		}
		count++
	}
	if count == maxIterationsFindHostPort {
		return -1, errors.New("Maximum iteration reached when finding host port, try creating new one again")
	}
	return hostPort, nil
}

func UpdateLXCState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	state := r.FormValue("state")

	if err = lxcModel.UpdateState(id, state); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if state == "failed" {
		errorMsg := r.FormValue("error_message")
		if err = lxcModel.UpdateErrorMessage(id, errorMsg); err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func UpdateLXCIp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	ip := r.FormValue("ip")

	if err = lxcModel.UpdateIpAddress(id, ip); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func DeleteLXC(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = lxcModel.DeleteLXC(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

func GetLXCsByLXDId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	lxcs, err := lxcModel.GetLXCsByLXDId(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, lxcs)
}

func getRandomPortNumber(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	port := min + randomizer.Intn(max-min)

	return port
}

func isLXCFormValid(lxcName, lxcImage, lxcPort string) bool {
	return len(lxcName) != 0 && len(lxcImage) != 0 && len(lxcPort) != 0
}

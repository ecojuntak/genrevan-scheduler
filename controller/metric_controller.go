package controller

import (
	"net/http"
	"strconv"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/gorilla/mux"
)

const (
	maxCounter = 100
)

var metricModel model.Metric

func UpdateMetric(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	metric, err := metricModel.GetMetricByLXDId(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	CPULoad, err := strconv.ParseFloat(r.FormValue("cpu"), 64)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	memoryLoad, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	newMetric := calculateNewState(CPULoad, memoryLoad, metric)
	metricModel.UpdateMetric(newMetric)

	RespondWithJSON(w, http.StatusOK, newMetric)
}

func calculateNewState(CPULoad float64, memoryLoad int, metric *model.Metric) model.Metric {
	CPUAverage := metric.CpuUsage
	memoryAverage := metric.MemoryUsage
	counter := metric.Counter

	if counter == maxCounter {
		counter--
	}

	CPUTotal := CPUAverage * float64(counter)
	memoryTotal := memoryAverage * counter
	counter = counter + 1

	newMetric := model.Metric{
		Id:          metric.Id,
		IdLxd:       metric.IdLxd,
		CpuUsage:    (CPUTotal + CPULoad) / float64(counter),
		MemoryUsage: (memoryTotal + memoryLoad) / counter,
		Counter:     counter,
	}

	return newMetric
}

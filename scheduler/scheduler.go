package scheduler

import "github.com/go-squads/genrevan-scheduler/model"

var metricModel model.Metric
var lxcModel model.Lxc

func GetLowestMetricLoad() (*model.Metric, error) {
	metrics, err := metricModel.GetMetricsBelowThreshold()

	if len(metrics) < 1 {
		metrics, err = metricModel.GetMetrics()
	}

	if err != nil {
		return nil, err
	}

	return &metrics[0], nil
}

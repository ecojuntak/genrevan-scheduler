package model_test

import (
	"strings"
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateMetric_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	metricId, err := metricModel.CreateMetric(1)
	assert.Equal(t, nil, err)

	metric, err := metricModel.GetMetric(*metricId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, metric.IdLxd)
}

func TestCreateMetric_ExpectedDuplicateLXDId(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	_, err = metricModel.CreateMetric(1)
	assert.Equal(t, nil, err)

	_, err = metricModel.CreateMetric(1)
	assert.True(t, strings.Contains(err.Error(), "duplicate key value violates unique constraint"))
}

func TestUpdateMetric_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	_, err = metricModel.CreateMetric(1)
	assert.Equal(t, nil, err)

	metric := model.Metric{
		IdLxd:       1,
		CpuUsage:    20.0003,
		MemoryUsage: 4096,
		Counter:     2,
	}

	err = metricModel.UpdateMetric(metric)
	assert.Equal(t, nil, err)
}

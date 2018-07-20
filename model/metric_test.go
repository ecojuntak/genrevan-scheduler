package model_test

import (
	"strings"
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

var metricModel model.Metric

func TestCreateMetric_ExpectedSuccess(t *testing.T) {
	setup()

	metricId, err := metricModel.CreateMetric(1)
	assert.Equal(t, nil, err)

	metric, err := metricModel.GetMetric(*metricId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, metric.IdLxd)
}

func TestCreateMetric_ExpectedDuplicateLXDId(t *testing.T) {
	setup()

	_, err := metricModel.CreateMetric(1)
	assert.Equal(t, nil, err)

	_, err = metricModel.CreateMetric(1)
	assert.True(t, strings.Contains(err.Error(), "duplicate key value violates unique constraint"))
}

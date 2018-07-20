package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterLXD_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	lxd, err := lxdModel.GetLXD("196.127.123.123")
	assert.Equal(t, nil, err)
	assert.Equal(t, "196.127.123.123", lxd.IpAddress)
}

func TestRegisterLXD_ExpectedMetricCreated(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	lxd, err := lxdModel.GetLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	metric, err := metricModel.GetMetricByLXDId(lxd.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, metric.IdLxd, lxd.Id)
}

func TestRegisterLXD_ExpectedDuplicated(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	err = lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)
}

func TestGetLXD_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxd, err := lxdModel.GetLXD("127.0.0.1")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lxd.Id)
}

func TestGetLXDs_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxds, err := lxdModel.GetLXDs()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, lxds)
}

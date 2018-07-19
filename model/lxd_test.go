package model_test

import (
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

var lxdModel model.Lxd

func TestRegisterLXD_ExpectedSuccess(t *testing.T) {
	setup()

	err := lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	lxd, err := lxdModel.GetLXD("196.127.123.123")
	assert.Equal(t, nil, err)
	assert.Equal(t, "196.127.123.123", lxd.IpAddress)
}

func TestRegisterLXD_ExpectedDuplicated(t *testing.T) {
	setup()

	err := lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)

	err = lxdModel.CreateLXD("196.127.123.123")
	assert.Equal(t, nil, err)
}

func TestGetLXD_ExpectedSuccess(t *testing.T) {
	setup()

	lxd, err := lxdModel.GetLXD("127.0.0.1")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lxd.Id)
}

func TestGetLXDs_ExpectedSuccess(t *testing.T) {
	setup()

	lxds, err := lxdModel.GetLXDs()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, lxds)
}

package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-squads/genrevan-scheduler/model"
)

var lxcModel model.Lxc

func setup() {
	model.SetupDatabase("testing")
}

func TestGetLXCs_ExpectedSuccess(t *testing.T) {
	setup()
	lxcs, err := lxcModel.GetLXCs()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, lxcs)
}

func TestGetLXC_ExpectedSuccess(t *testing.T) {
	setup()
	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lxc.Id)
}

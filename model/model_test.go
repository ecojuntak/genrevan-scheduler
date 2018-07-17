package model_test

import (
	"errors"
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

func TestSetupDatabase_ExpectedConnect(t *testing.T) {
	err := model.SetupDatabase("testing")
	assert.Equal(t, nil, err)
}

func TestSetupDatabase_ExpectedWrongEnvironment(t *testing.T) {
	err := model.SetupDatabase("")
	assert.Equal(t, errors.New("Environment not match"), err)
}

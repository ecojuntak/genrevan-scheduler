package model_test

import (
	"errors"
	"testing"

	"github.com/go-squads/genrevan-scheduler/migration"
	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

var lxcModel model.Lxc
var lxdModel model.Lxd
var metricModel model.Metric

func TestSetupDatabase_ExpectedConnect(t *testing.T) {
	err := model.SetupDatabase("testing")
	assert.Equal(t, nil, err)
}

func TestSetupDatabase_ExpectedWrongEnvironment(t *testing.T) {
	err := model.SetupDatabase("")
	assert.Equal(t, errors.New("Environment not match"), err)
}

func setup() error {
	model.SetupDatabase("testing")
	err := migration.RunMigration()
	if err != nil {
		return err
	}

	err = migration.RunSeeder()

	if err != nil {
		return err
	}

	return nil
}

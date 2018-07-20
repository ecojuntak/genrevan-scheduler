package scheduler_test

import (
	"testing"

	"github.com/go-squads/genrevan-scheduler/migration"
	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/scheduler"
	"github.com/stretchr/testify/assert"
)

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

func TestScheduler_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	metric, err := scheduler.GetLowestMetricLoad()
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, metric.IdLxd)
}

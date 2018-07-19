package migration_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-squads/genrevan-scheduler/migration"
	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/util"
)

var basepath = util.GetRootFolderPath()

func TestRunMigration_ExpectedSuccess(t *testing.T) {
	err := migration.RunMigration()
	assert.Equal(t, nil, err)
}

func TestGetStringFromFile_ExpectedFound(t *testing.T) {
	_, err := migration.GetStringFromFile(basepath + "migration/schema.sql")
	assert.Equal(t, nil, err)
}

func TestGetStringFromFile_ExpectedFileNotFound(t *testing.T) {
	_, err := migration.GetStringFromFile("schema")
	assert.Equal(t, errors.New("File not found"), err)
}

func TestRunSeeder_ExpectedSuccess(t *testing.T) {
	err := migration.RunMigration()
	assert.Equal(t, nil, err)
	err = migration.RunSeeder()
	assert.Equal(t, nil, err)
}

func TestRunSeeder_ExpectedDataCreated(t *testing.T) {
	err := migration.RunMigration()
	assert.Equal(t, nil, err)
	err = migration.RunSeeder()
	assert.Equal(t, nil, err)

	var lxcModel model.Lxc
	lxc, err := lxcModel.GetLXC(1)

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lxc.Id)
}

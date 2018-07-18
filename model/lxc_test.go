package model_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-squads/genrevan-scheduler/migration"
	"github.com/go-squads/genrevan-scheduler/model"
)

var lxcModel model.Lxc

func setup() {
	model.SetupDatabase("testing")
	migration.RunMigration()
	migration.RunSeeder()
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

func TestCreateLXC_ExpectedDataCreated(t *testing.T) {
	setup()

	lxc := model.Lxc{
		Name:  "GO-PAY System Configuration",
		Image: "xenial64",
	}

	newLxcId, err := lxcModel.CreateLXC(lxc)
	assert.Equal(t, nil, err)

	newLxc, err := lxcModel.GetLXC(*newLxcId)
	assert.Equal(t, "GO-PAY System Configuration", newLxc.Name)
	assert.Equal(t, "xenial64", newLxc.Image)
	assert.Equal(t, "pending", newLxc.Status)
}

func TestDeleteLXC_ExpectedDataDeleted(t *testing.T) {
	setup()

	err := lxcModel.DeleteLXC(3)
	assert.Equal(t, nil, err)

	lxc, err := lxcModel.GetLXC(3)
	assert.Empty(t, lxc)
	assert.NotEqual(t, nil, err)
}

func TestUpdateLXCIpAddress_ExpectedSuccess(t *testing.T) {
	setup()

	err := lxcModel.UpdateIpAddress(1, "123.123.123.123")
	assert.Equal(t, nil, err)

	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "123.123.123.123", lxc.IpAddress.String)
}

func TestUpdateLXCIpAddress_ExpectedErrorDuplicateIp(t *testing.T) {
	setup()

	err := lxcModel.UpdateIpAddress(1, "127.0.0.1")
	assert.Equal(t, nil, err)
	err = lxcModel.UpdateIpAddress(2, "127.0.0.1")
	assert.True(t, strings.Contains(err.Error(), "duplicate key value violates unique constraint"))
}

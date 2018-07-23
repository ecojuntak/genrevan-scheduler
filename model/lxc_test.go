package model_test

import (
	"strings"
	"testing"

	"github.com/guregu/null"

	"github.com/stretchr/testify/assert"

	"github.com/go-squads/genrevan-scheduler/model"
)

func TestGetLXCs_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxcs, err := lxcModel.GetLXCs()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, lxcs)
}

func TestGetLXC_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lxc.Id)
}

func TestGetLXCsByLXDId_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxcs, err := lxcModel.GetLXCsByLXDId(1)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, lxcs)
	assert.Equal(t, 2, len(lxcs))
}

func TestCreateLXC_ExpectedDataCreated(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	lxc := model.Lxc{
		Name:  "GO-PAY System Configuration",
		Image: "xenial64",
		LxdId: null.NewInt(1, true),
	}

	newLxcId, err := lxcModel.CreateLXC(lxc)
	assert.Equal(t, nil, err)

	newLxc, err := lxcModel.GetLXC(*newLxcId)
	assert.Equal(t, nil, err)
	assert.Equal(t, "GO-PAY System Configuration", newLxc.Name)
	assert.Equal(t, "xenial64", newLxc.Image)
	assert.Equal(t, "pending", newLxc.Status)
}

func TestDeleteLXC_ExpectedDataDeleted(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxcModel.DeleteLXC(3)
	assert.Equal(t, nil, err)

	lxc, err := lxcModel.GetLXC(3)
	assert.Empty(t, lxc)
	assert.NotEqual(t, nil, err)
}

func TestUpdateLXCIpAddress_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxcModel.UpdateIpAddress(1, "123.123.123.123")
	assert.Equal(t, nil, err)

	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "123.123.123.123", lxc.IpAddress.String)
}

func TestUpdateLXCIpAddress_ExpectedErrorDuplicateIp(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxcModel.UpdateIpAddress(1, "127.0.0.1")
	assert.Equal(t, nil, err)
	err = lxcModel.UpdateIpAddress(2, "127.0.0.1")
	assert.True(t, strings.Contains(err.Error(), "duplicate key value violates unique constraint"))
}

func TestUpdateLXCState_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxcModel.UpdateState(1, "running")
	assert.Equal(t, nil, err)
	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "running", lxc.Status)
}

func TestUpdateLXCLXDId_ExpectedSuccess(t *testing.T) {
	err := setup()
	assert.Equal(t, nil, err)

	err = lxcModel.UpdateLxdId(1, 1)
	assert.Equal(t, nil, err)
	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), lxc.LxdId.Int64)
}

package router_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/guregu/null"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

func TestGetLXCsRouter_ExpecetedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxc", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxcs := []model.Lxc{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxcs)
	assert.Equal(t, nil, err)

	assert.Equal(t, 3, len(lxcs))
}

func TestGetLXCRouter_ExpectedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxc/1", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxc := model.Lxc{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxc)

	assert.Equal(t, 1, lxc.Id)
}

func TestCreateLXC_ExpectedStatusCreated(t *testing.T) {
	data := url.Values{}
	data.Set("name", "foo")
	data.Set("image", "xenial64")
	data.Set("containerPort", "80")

	req, err := http.NewRequest("POST", "/lxc", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestCreateLXC_ExpectedStatusBadRequest(t *testing.T) {
	data := url.Values{}
	data.Set("name", "")
	data.Set("image", "")

	req, err := http.NewRequest("POST", "/lxc", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestUpdateLXCState_ExpectedStatusSucccess(t *testing.T) {
	data := url.Values{}
	data.Set("state", "running")

	req, err := http.NewRequest("PATCH", "/lxc/1/state", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusNoContent, response.Code)

	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "running", lxc.Status)
}

func TestUpdateLXCIp_ExpectedStatusSucccess(t *testing.T) {
	data := url.Values{}
	data.Set("ip", "192.168.1.1")

	req, err := http.NewRequest("PATCH", "/lxc/1/ip", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusNoContent, response.Code)

	lxc, err := lxcModel.GetLXC(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, null.NewString("192.168.1.1", true), lxc.IpAddress)
}

func TestGetLXCsByLXDId_ExpectedSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxc/lxd/1", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxcs := []model.Lxc{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxcs)

	assert.Equal(t, 2, len(lxcs))
}

func TestDeleteLXC_ExpectedStatusSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/lxc/1", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusNoContent, response.Code)

	_, err = lxcModel.GetLXC(1)
	assert.NotEqual(t, nil, err)
}

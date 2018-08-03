package router_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

func TestGetLXDsRouter_ExpectedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxd", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxds := []model.Lxd{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxds)

	assert.Equal(t, 3, len(lxds))
}

func TestGetLXDRouter_ExpectedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxd/1", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxc := model.Lxc{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxc)

	assert.Equal(t, 1, lxc.Id)
}

func TestRegisterLXD_ExpectedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxd/register", nil)
	req.RemoteAddr = "127.0.0.4:36496"
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusCreated, response.Code)

	dataResponse := make(map[string]int)
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &dataResponse)

	assert.Equal(t, 4, dataResponse["id"])
}

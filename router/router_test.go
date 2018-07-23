package router_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-squads/genrevan-scheduler/migration"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/router"
	"github.com/stretchr/testify/assert"
)

const env = "testing"

func init() {
	model.SetupDatabase(env)
	migration.RunMigration(env)
	migration.RunSeeder(env)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := router.SetupRouter()
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	return response
}

func TestGetLXCsRouter_ExpecetedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxc", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)

	lxcs := []model.Lxc{}
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &lxcs)

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

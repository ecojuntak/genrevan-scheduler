package router_test

import (
	"net/http"
	"net/http/httptest"
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

func TestGetLXCRouter_ExpecetedStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/lxc", nil)
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)
}

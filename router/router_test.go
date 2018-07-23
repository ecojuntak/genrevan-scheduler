package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-squads/genrevan-scheduler/migration"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/router"
)

const env = "testing"

var lxcModel model.Lxc

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

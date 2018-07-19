package util_test

import (
	"strings"
	"testing"

	"github.com/go-squads/genrevan-scheduler/util"
	"github.com/stretchr/testify/assert"
)

func TestGetRootPath_ExpectedSuccess(t *testing.T) {
	path := util.GetRootFolderPath()
	splittedPath := strings.Split(path, "/")
	assert.Equal(t, "genrevan-scheduler", splittedPath[len(splittedPath)-2])
}

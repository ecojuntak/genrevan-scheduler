package config_test

import (
	"errors"
	"testing"

	"github.com/go-squads/genrevan-scheduler/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig_ExpectedFileFound(t *testing.T) {
	var conf config.Conf
	_, err := conf.GetConfig("testing.example.yaml")

	assert.Equal(t, nil, err)
}

func TestConfig_ExpectedFileNotFound(t *testing.T) {
	var conf config.Conf
	_, err := conf.GetConfig("testing.json")

	assert.Equal(t, errors.New("File not found"), err)
}

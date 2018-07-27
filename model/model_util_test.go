package model_test

import (
	"testing"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateTooLongName_ExpectedFail(t *testing.T) {
	s := "A01234567801234567801234567801234567801234567801234567801234567801234567"

	err := model.ValidateLXCName(s)
	assert.NotEqual(t, nil, err)
}

func TestValidateNameStartHypen_ExpectedFail(t *testing.T) {
	s := "-invalid"

	err := model.ValidateLXCName(s)
	assert.NotEqual(t, nil, err)
}

func TestValidateNameStartNumber_ExpectedFail(t *testing.T) {
	s := "04invalid"

	err := model.ValidateLXCName(s)
	assert.NotEqual(t, nil, err)
}

func TestValidateNameEndHypen_ExpectedFail(t *testing.T) {
	s := "invalid-"

	err := model.ValidateLXCName(s)
	assert.NotEqual(t, nil, err)
}

func TestValidateName_ExpectedSuccess(t *testing.T) {
	s := "valid-ubuntu64"

	err := model.ValidateLXCName(s)
	assert.Equal(t, nil, err)
}

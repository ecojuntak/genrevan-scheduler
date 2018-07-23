package router_test

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMetricRouter_ExpectedSuccess(t *testing.T) {
	data := url.Values{}
	data.Set("cpu", "12.34")
	data.Set("memory", "1233")

	req, err := http.NewRequest("PUT", "/metric/2", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)
}

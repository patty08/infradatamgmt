package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetActionStdout(t *testing.T) {
	service := ServiceStdout{}

	data := map[string]string {
		"monitor" : "enabled",
		"logging" : "enabled",
	}
	err := service.GetAction("create", data)
	assert.Nil(t, err)

	err = service.GetAction("", data)
	assert.NotNil(t, err)
}

package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetActionLogging(t *testing.T) {
	service := ServiceLogging{}

	err := service.GetAction("create", map[string]string {})
	assert.Nil(t, err)

	//TODO write unit test
}

package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestSetActionMetrics(t *testing.T) {
	service := ServiceMetrics{}

	err := service.SetAction("create", map[string]string {})
	assert.Nil(t, err)

	//TODO write unit test
}

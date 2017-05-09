package rooter

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sebastienmusso/infradatamgmt/agent"
)

func TestProcess(t *testing.T)  {
	err := process(nil)
	assert.NotNil(t, err)

	info := &agent.InfoIN{
		Action : "create",
		Services : []string{},
		Data : map[string]string{},
	}

	err = process(info)
	assert.Nil(t, err)

	info.Data["client"] = "docker"
	err = process(info)
	assert.Nil(t, err)
}
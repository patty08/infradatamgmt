package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
)

func TestSetActionMetrics(t *testing.T) {
	service := ServiceMetrics{}

	err := service.GetAction("create", map[string]string {}, nil)
	assert.Nil(t, err)
}
func TestFormatMetricHostName(t *testing.T)  {

	ip := "127.0.0.1"
	out := formatMetricHostName(ip)
	assert.Equal(t,"hosts: [\"127.0.0.1\"]",out,"Error Reformat Host name")
	ip = ""
	out =formatMetricHostName(ip)
	assert.Equal(t,"hosts: [\"\"]",out,"Error Empty host")

}

func TestDetachMetricsConfiguration(t *testing.T)  {

	data := map[string]string{}

	data["application_type"]= "test"
	data["id"]= "file"

	_, err := os.Create(data["application_type"]+"_"+data["id"]+".yml")
	assert.NotNil(t,err)

	detachMetricConfiguration(data)
	_,err = os.Open(data["application_type"]+"_"+data["id"]+".yml")
	assert.NotNil(t,err)

}
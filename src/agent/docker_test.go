package agent

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestConnectDocker(t *testing.T) {
	_ , err := connectDocker("unix:///var/run/docker.sock")
	assert.Nil(t, err)
	_ , err = connectDocker("")
	assert.NotNil(t, err)
}

func TestParseDockerEvent(t *testing.T) {
	listener := make(chan *InfoIN, 1)
	eventData := events.Message{
		ID :"3034",
		From :"ubuntu",
		Action :"create",
		Actor : events.Actor{
			ID : "3034",
			Attributes : map[string]string {
				"image": "ubuntu",
				"logging": "enabled",
				"monitor": "enabled",
				"name": "nAmE",
			},
		},
		Time : 1493883345,
	}

	parseDockerEvent(eventData, listener)
	out := <-listener

	assert.Equal(t, "create", out.Action, "Error parse Action")
	assert.Equal(t, []string {"logging"}, out.Services, "Error parse Services")
	assert.Equal(t, "ubuntu", out.Data["image"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["logging"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["monitor"], "Error parse Data")
	assert.Equal(t, "nAmE", out.Data["name"], "Error parse Data")
	assert.Equal(t, "1493883345", out.Data["timestamp"], "Error parse Data")
	assert.Equal(t, time.Unix(1493883345, 0).String(), out.Data["time"], "Error parse Date in Data")

	eventData = events.Message{
		ID :"3034",
		Action :"die",
		Actor : events.Actor{
			ID : "3034",
			Attributes : map[string]string {
				"logging": "disabled",
			},
		},
	}

	parseDockerEvent(eventData, listener)
	out = <-listener

	assert.Equal(t, "stop", out.Action, "Error parse Action")
	assert.Equal(t, []string {}, out.Services, "Error parse Services")
	assert.Equal(t, "disabled", out.Data["logging"], "Error parse Data")
}

func TestParseDockerEventAction(t *testing.T) {
	listener := make(chan *InfoIN, 1)

	eventData := events.Message{Action :"destroy",}
	parseDockerEvent(eventData, listener)
	out := <-listener

	assert.Equal(t, "delete", out.Action, "Error parse Action")

	eventData = events.Message{Action :"unpause",}
	parseDockerEvent(eventData, listener)
	out = <-listener

	assert.Equal(t, "start", out.Action, "Error parse Action")

	eventData = events.Message{Action :"pause",}
	parseDockerEvent(eventData, listener)
	out = <-listener

	assert.Equal(t, "stop", out.Action, "Error parse Action")
}

func TestAddDockerListener(t *testing.T) {
	listener := make(chan *InfoIN, 1)
	client , err := connectDocker("unix:///var/run/docker.sock")
	assert.Nil(t, err)

	go addDockerListener(client, listener)
	cfg := &container.Config{
		Image: "busybox",
		Labels: map[string]string {
			"monitor" : "enabled",
			"logging" : "enabled",
		},
	}
	client.ContainerCreate(context.Background(), cfg, nil, nil, "unit_test")

	out := <-listener
	assert.Equal(t, "create", out.Action, "Error parse Action")
	assert.Equal(t, []string {"logging"}, out.Services, "Error parse Services")
	assert.Equal(t, "busybox", out.Data["image"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["logging"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["monitor"], "Error parse Data")

	client.ContainerRemove(context.Background(), out.Data["id"], types.ContainerRemoveOptions{})

	out = <-listener
	assert.Equal(t, "delete", out.Action, "Error parse Action")

}

func TestAddEventListenerError(t *testing.T) {
	listener := make(chan *InfoIN, 1)

	agent := AgentDocker{}
	err := agent.AddEventListener(listener, "")
	assert.NotNil(t, err)
}

func TestAddEventListener(t *testing.T) {
	listener := make(chan *InfoIN, 1)
	client , err := connectDocker("unix:///var/run/docker.sock")
	assert.Nil(t, err)

	agent := AgentDocker{}
	go agent.AddEventListener(listener, "unix:///var/run/docker.sock")

	cfg := &container.Config{
		Image: "busybox",
		Labels: map[string]string {
			"monitor" : "enabled",
			"logging" : "enabled",
		},
	}
	client.ContainerCreate(context.Background(), cfg, nil, nil, "unit_test")

	out := <-listener
	assert.Equal(t, "create", out.Action, "Error parse Action")
	assert.Equal(t, []string {"logging"}, out.Services, "Error parse Services")
	assert.Equal(t, "busybox", out.Data["image"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["logging"], "Error parse Data")
	assert.Equal(t, "enabled", out.Data["monitor"], "Error parse Data")

	client.ContainerRemove(context.Background(), out.Data["id"], types.ContainerRemoveOptions{})

	out = <-listener
	assert.Equal(t, "delete", out.Action, "Error parse Action")
}
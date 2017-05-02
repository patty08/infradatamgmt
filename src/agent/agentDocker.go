package agent

import (
	"github.com/fsouza/go-dockerclient"
	"fmt"
)

// Struct for strategy module
type AgentDocker struct {}

// Add Event Listener in Docker client
// IN : main chan for send InfoIn event
func (AgentDocker) AddEventListener(main chan *InfoIN, who string) error {
	endpoint := who
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return fmt.Errorf("Unable to start Docker EventListener :\n- %s", err)
	}

	listener := make(chan *docker.APIEvents)
	if err := client.AddEventListener(listener); err != nil {
		return fmt.Errorf("Unable to start Docker EventListener :\n- %s", err)
	}

	for {
		f := <-listener
		go parseDockerEvent(client, f)
	}
}

func parseDockerEvent(client *docker.Client, event *docker.APIEvents) {

	// TODO, chose event, parse in InfoIn format

	fmt.Print(event.Type + " " + event.Action )

}
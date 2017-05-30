package agent

import (
   "fmt"
   "github.com/docker/docker/api/types"
   "github.com/docker/docker/api/types/events"
   "github.com/docker/docker/api/types/filters"
   "github.com/docker/docker/client"
   "golang.org/x/net/context"
   "strconv"
   "time"

)

// Label for activate monitoring.
const label_monitoring string = "monitor"

// Struct for strategy module
type AgentDocker struct {}

// Add Event Listener in Docker client.
// IN : main chan for send InfoIn event.
func (AgentDocker) AddEventListener(main chan *InfoIN, who string) error {
	client, err := connectDocker(who)
	if err != nil {
		return fmt.Errorf("Unable to start Docker EventListener :\n- %s", err)
	}

   addDockerListener(client, main)

   return nil
}

// Connect agent to docker API.
func connectDocker(who string) (*client.Client, error) {
	client, err := client.NewClient(who, "1.25", nil, nil)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Successfully connected to docker socket")
	}
	return client, nil
}

// Start event listener on docker client.
func addDockerListener (client *client.Client, main chan *InfoIN) {
	fmt.Println("Successfully start Event Listener")

	// listen to event values to run an action
	f := filters.NewArgs()
	f.Add("event", "start")
	f.Add("event", "create")
	f.Add("event", "die")
	f.Add("event", "destroy")
	f.Add("type", "container")
	f.Add("label", label_monitoring+"=enabled")
	options := types.EventsOptions{Filters: f}

	ctx, cancel := context.WithCancel(context.Background())
	eventsChan, errChan := client.Events(ctx, options)

	go func(){
		for event := range eventsChan {
			go processDockerEvent(event, main, client)
		}

	}()

	if err := <-errChan; err != nil {
		fmt.Println("Event monitor throw this error: ", err)
	}

	defer cancel()
}

// Process docker event information for rooter.
func processDockerEvent(event events.Message, main chan *InfoIN, client *client.Client)  {
	data := map[string]string{}

	// Inspect Container for more information
	inspect, err := client.ContainerInspect(context.Background(), event.ID)
	if err == nil {
		data["ip"] = string(inspect.NetworkSettings.IPAddress)
	}

	parseDockerEvent(event, data, main)
}

// Parse docker event information for rooter.
func parseDockerEvent(event events.Message, inspect map[string]string, main chan *InfoIN)  {
	infos := &InfoIN{}
	infos.Action = event.Action
	if infos.Action == "die" || infos.Action == "pause" {
		infos.Action = "stop"
	} else if infos.Action == "unpause" {
		infos.Action = "start"
	} else if infos.Action == "destroy" {
		infos.Action = "delete"
	}

	infos.Services = []string{}
	infos.Data = map[string]string{}

	infos.Data["id"] = event.ID
	infos.Data["action"] = event.Action
	infos.Data["type"] = event.Type
	infos.Data["timestamp"] = strconv.FormatInt(event.Time, 10)
	infos.Data["time"] = time.Unix(event.Time, 0).String()

	for k,v := range inspect {
		infos.Data[k] = v
	}

	for k,v := range event.Actor.Attributes {
		if v == "enabled" && k != label_monitoring {
			infos.Services = append(infos.Services, k)
		}
		infos.Data[k] = v
	}

	main <- infos
}
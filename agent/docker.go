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
   "os"
)

// Label for activate monitoring.
const label_monitoring string = "monitor"
const dirCustom string = "./rooter/configuration/metricbeat/conf/custom/"

// Struct for strategy module
type AgentDocker struct {}

// Add Event Listener in Docker client.
// IN : main chan for send InfoIn event.
func (AgentDocker) AddEventListener(main chan *InfoIN, who string) error {
	client, err := connectDocker(who);
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
   f.Add("event", "die")
   // TODO: f.Add("event", "destroy") , pouvoir detacher le container (supprimer fichier de configuration, etc)
   f.Add("type", "container")
   f.Add("label", label_monitoring+"=enabled")
   options := types.EventsOptions{Filters: f}

   ctx, cancel := context.WithCancel(context.Background())
   eventsChan, errChan := client.Events(ctx, options)

   go func(){
	  for event := range eventsChan {
		 go parseDockerEvent(event, main, client)
	  }

   }()

   if err := <-errChan; err != nil {
	  fmt.Println("Event monitor throw this error: ", err)
   }

   defer cancel()
}

// Parse docker envent information for rooter.
func parseDockerEvent(event events.Message, main chan *InfoIN, client *client.Client)  {
	infos := &InfoIN{}
	infos.Action = event.Action
	if infos.Action == "die" || infos.Action == "pause" {
		infos.Action = "stop"
	   	println("TODO delete configuration")
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
   infos.Data["ip"] = string(getIpContainer(client, event.ID))
   for k,v := range event.Actor.Attributes {
		if v == "enabled" && k != label_monitoring {
			infos.Services = append(infos.Services, k)
		}
		infos.Data[k] = v
	}

	main <- infos
}
// get id's container in instance and return as a string.
// client is current running container.
// containerID is the id of the container.
func getIpContainer(client *client.Client, containerID  string) string{
   inspect_result, err := client.ContainerInspect(context.Background(), containerID)
   if err != nil {
	  println("error")
   }
   ipHost := inspect_result.NetworkSettings.IPAddress
   return ipHost
}

// TODO: Function to remove container's configuration file
func detachContainer(main *InfoIN){
   err:= os.Remove(dirCustom +main.Data["application_type"]+"_"+main.Data["id"]+".yml")
   if err != nil {
	  fmt.Println(err)
	  return
   }
	println("Detached..")
}
package client

import (
    "github.com/sebastienmusso/infradatamgmt/agent"
    "github.com/sebastienmusso/infradatamgmt/service"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"fmt"
	"time"
)

// Client structure.
type ClientDocker struct {}

// Function choose the action to set for the agent services.
// Services is the list of services to activate.
// Data is all informations data send by the container.
func (ClientDocker) SetAction(info *agent.InfoIN) error {
   for k := 0 ; k <= len(info.Services)-1; k++{
	  switch info.Services[k]{
	  case "stdout":
		 {
			 service := sServiceOut{service.ServiceStdout{}}
			 go service.aServiceOut.GetAction(info.Action, info.Data, nil)
		 }
	  case "logging":
		 {
			 l := make(chan *service.ClientIN)
			 service := sServiceOut{service.ServiceLogging{}}
			 go service.aServiceOut.GetAction(info.Action, info.Data, l)
			 go listener(l)
		 }
	  case "metric":
		 {
			 service := sServiceOut{service.ServiceMetrics{}}
			 go service.aServiceOut.GetAction(info.Action, info.Data, nil)
			 // setConfigServices(info.Data["image"], info.Data["application_type"], info.Data["id"])
		 }
	  }
   }
   return nil
}

func listener(c chan *service.ClientIN) {
	for {
		i := <-c

		switch i.Action {
		case "start":
			cfg, hostCfg := containerConfig(i.Data)
			startContainer(i.Data["name"], cfg, hostCfg)
		case "stop":

		case "end":
			close(c)
			return
		default:
		}

		time.Sleep(time.Second * 1)
	}
}

func containerConfig(data map[string]string) (*container.Config, *container.HostConfig) {
	cfg := &container.Config{
		Image: data["image"],
		Labels: map[string]string {
			"maintainer" : "surikator",
			"associate-name" : data["who_name"],
			"associate-id" : data["who_id"],
		},
	}

	if data["volume"] {
		cfg.Volumes = map[string]struct{} {
			data["volume_src"] : {data["volume_container"]},
		}
	}

	hostCfg := &container.HostConfig{}
	if data["volume_from"] {
		hostCfg.VolumesFrom = []string{data["who_name"]}
	}
	return cfg, hostCfg
}

func startContainer(name string, cfg *container.Config, hostCfg *container.HostConfig) error {
	client, err := connectDocker("unix:///var/run/docker.sock")
	if err != nil {
		return fmt.Errorf("Unable to start Docker client :\n- %s", err)
	}

	r, err := client.ContainerCreate(context.Background(), cfg, hostCfg, nil, name)
	if err != nil {
		fmt.Println("Could not create container", err)
		return err
	}
	err = client.ContainerStart(context.Background(), r.ID, types.ContainerStartOptions{})
	if err != nil {
		fmt.Println("Cannot start container", err)
		return err
	}

	return nil
}

// Connect to docker API.
func connectDocker(who string) (*client.Client, error) {
	client, err := client.NewClient(who, "1.25", nil, nil)
	if err != nil {
		fmt.Println("Error connection to docker socket")
		return nil, err
	}
	return client, nil
}

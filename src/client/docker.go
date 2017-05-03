package client

import (
   "agent"
   "service"
   "fmt"
)

type ClientDocker struct {}

// function choose the action to set for the agent services
// services is the list of services to activate
// data is all informations data send by the container

func (ClientDocker) SetAction(info *agent.InfoIN) error {
   for k := 0 ; k <= len(info.Services)-1; k++{
	  switch info.Services[k]{
	  case "stdout":
		 {
			agent := sServiceOut{service.ServiceStdout{}}
			go agent.aServiceOut.SetAction(info.Action, info.Data)
			fmt.Println("ici")
		 }
	  case "logging":
		 {
			agent := sServiceOut{service.ServiceLogging{}}
			go agent.aServiceOut.SetAction(info.Action, info.Data)
		 }
	  case "metric":
		 {
			agent := sServiceOut{service.ServiceMetrics{}}
			go agent.aServiceOut.SetAction(info.Action, info.Data)
		 }
	  }
   }

   return nil
}

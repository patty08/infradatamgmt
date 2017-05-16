package client

import (
   "github.com/sebastienmusso/infradatamgmt/agent"
   "github.com/sebastienmusso/infradatamgmt/service"
)

//Structure for Docker client.
type ClientDocker struct {}

// Function choose the action to set for the agent services.
// Services is the list of services to activate.
// Data is all informations data send by the container.
func (ClientDocker) SetAction(info *agent.InfoIN) error {
   for k := 0 ; k <= len(info.Services)-1; k++{
	  switch info.Services[k]{
	  case "stdout":
		 {
			agent := sServiceOut{service.ServiceStdout{}}
			go agent.aServiceOut.SetAction(info.Action, info.Data)
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

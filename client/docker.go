package client

import (
   "github.com/sebastienmusso/infradatamgmt/agent"
   "github.com/sebastienmusso/infradatamgmt/service"
)

<<<<<<< HEAD
//Structure for Docker client.
type ClientDocker struct {}

=======

type ClientDocker struct {}

>>>>>>> dee82e212a07b920d050360d0c910f991e60807c
// Function choose the action to set for the agent services.
// Services is the list of services to activate.
// Data is all informations data send by the container.
func (ClientDocker) SetAction(info *agent.InfoIN) error {
   for k := 0 ; k <= len(info.Services)-1; k++{
	  switch info.Services[k]{
	  case "stdout":
		 {
			service := sServiceOut{service.ServiceStdout{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
		 }
	  case "logging":
		 {
			service := sServiceOut{service.ServiceLogging{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
		 }
	  case "metric":
		 {
			service := sServiceOut{service.ServiceMetrics{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
			// setConfigServices(info.Data["image"], info.Data["application_type"], info.Data["id"])
		 }
	  }
   }
   return nil
}
